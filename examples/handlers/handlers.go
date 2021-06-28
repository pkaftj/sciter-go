package main

import (
	"flag"
	"log"
	"path/filepath"
	
	"github.com/pkaftj/sciter-go"
	"github.com/pkaftj/sciter-go/rice"
	"github.com/pkaftj/sciter-go/window"
)

func setEventHandler(w *window.Window) {
	w.DefineFunction("kkk", func(args ...*sciter.Value) *sciter.Value {
		log.Println("kkk called!")
		rval := sciter.NewValue()
		rval.Set("num", sciter.NewValue(1000))
		rval.Set("str", sciter.NewValue("a string"))
		var fn sciter.NativeFunctor
		fn = func(args ...*sciter.Value) *sciter.Value {
			log.Printf("args[%d]:%v\n", len(args), args)
			return sciter.NewValue("native functor called")
		}
		rval.Set("f", sciter.NewValue(fn))
		return rval
	})
	w.DefineFunction("sumall", func(args ...*sciter.Value) *sciter.Value {
		sum := 0
		for i := 0; i < len(args); i++ {
			sum += args[i].Int()
		}
		return sciter.NewValue(sum)
	})
	w.DefineFunction("gprintln", func(args ...*sciter.Value) *sciter.Value {
		print("\t->")
		for _, arg := range args {
			print(arg.String(), "\t")
		}
		println()
		return sciter.NullValue()
	})
}

func setCallbackHandlers(w *window.Window) {
	h := &sciter.CallbackHandler{
		OnDataLoaded: func(ld *sciter.ScnDataLoaded) int {
			log.Println("loaded", ld.Uri())
			return sciter.LOAD_OK
		},
		OnLoadData: func(ld *sciter.ScnLoadData) int {
			log.Println("loading", ld.Uri())
			return sciter.LOAD_OK
		},
	}
	w.SetCallback(h)
}

func setElementHandlers(root *sciter.Element) {
	log.Println("setElementHandlers")
	el, err := root.SelectFirst("#native")
	if err == nil {
		log.Println("SelectFirst:", el)
		el.OnClick(func() {
			log.Println("native handler called")
		})
	}
	root.DefineMethod("mcall", func(args ...*sciter.Value) *sciter.Value {
		print("->method args:")
		for _, arg := range args {
			print(arg.String(), "\t")
		}
		println()
		return sciter.NullValue()
	})
}

func testCall(w *window.Window) {
	// test sciter call
	v, err := w.Call("gFunc", sciter.NewValue("kkk"), sciter.NewValue(555))
	if err != nil {
		log.Println("sciter call failed:", err)
	} else {
		log.Println("sciter call successfully:", v.String())
	}

	// test method call
	root, err := w.GetRootElement()
	if err != nil {
		log.Fatalf("get root element failed: %s", err.Error())
	}
	v, err = root.CallMethod("mfn", sciter.NewValue("method call"), sciter.NewValue(10300))
	if err != nil {
		log.Println("method call failed:", err)
	} else {
		log.Println("method call successfully:", v.String())
	}
	// test function call
	v, err = root.CallFunction("gFunc", sciter.NewValue("function call"), sciter.NewValue(10300))
	if err != nil {
		log.Println("function call failed:", err)
	} else {
		log.Println("function call successfully:", v.String())
	}

	v, err = root.CallFunction("gFuncEmpty")
	if err != nil {
		log.Println("function call gFuncEmpty() failed:", err)
	} else {
		log.Println("function call gFuncEmpty() successfully:", v.String())
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("no html file found")
	}
	// create window
	w, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sciter Version: %X %X\n", sciter.Version(true), sciter.Version(false))
	// resource packing
	rice.HandleDataLoad(w.Sciter)
	// enable debug
	ok := w.SetOption(sciter.SCITER_SET_DEBUG_MODE, 1)
	if !ok {
		log.Println("set debug mode failed")
	}
	// load file
	fullpath, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if err = w.LoadFile(fullpath); err != nil {
		log.Println("LoadFile error:", err.Error())
		return
	}
	root, err := w.GetRootElement()
	if err != nil {
		log.Fatalf("get root element failed: %s", err.Error())
	}
	setElementHandlers(root)
	// set handlers
	setEventHandler(w)
	setCallbackHandlers(w)
	testCall(w)
	w.Show()
	w.Run()
}

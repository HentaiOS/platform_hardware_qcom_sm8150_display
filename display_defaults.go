package display

import (
        "android/soong/android"
        "android/soong/cc"
)

func init() {
	android.RegisterModuleType("display_go_defaults", display_DefaultsFactory)
}

func display_DefaultsFactory() (android.Module) {
	module := cc.DefaultsFactory()
	android.AddLoadHook(module, display_Defaults)
	return module
}

func display_Defaults(ctx android.LoadHookContext) {
	type props struct {
		Enabled *bool
	}
	p := &props{}
	p.Enabled = display_globalDefaults(ctx)
	ctx.AppendProperties(p)
}

func display_globalDefaults(ctx android.BaseContext) (*bool) {
	var module_enabled *bool

	return module_enabled
}


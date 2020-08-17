package generator

type astEditorOption = func(editor *astEditor)

// addBeforeEditHook add the given file hook to the editor before edit hooks.
func addBeforeEditHook(hook fileHook) astEditorOption {
	return func(editor *astEditor) {
		editor.beforeEditHooks = append(editor.beforeEditHooks, hook)
	}
}

// addNodeHook add the given node hook to the editor node hooks.
func addNodeHook(hook nodeHook) astEditorOption {
	return func(editor *astEditor) {
		editor.nodeHooks = append(editor.nodeHooks, hook)
	}
}

// addAfterEditHook add the given file hook to the editor after edit hooks.
func addAfterEditHook(hook fileHook) astEditorOption {
	return func(editor *astEditor) {
		editor.afterEditHooks = append(editor.afterEditHooks, hook)
	}
}

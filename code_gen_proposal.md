
Proposal
================================================================================

Write a code gen tool for wrapping a specific function signature in a struct,
that implements an interface which `labassistant` uses to create and run new
`experiments`.

Status
--------------------------------------------------------------------------------

Unfinished draft.

Pros
--------------------------------------------------------------------------------

- Not having to do dangerous `reflect` magic and be able to use strongly typed code.
- No `reflect` magic means less code to run, therefor speed boost?

Cons
--------------------------------------------------------------------------------

- Yet another tool to run.
- More code the user has to manage himself.

Design
--------------------------------------------------------------------------------

Run `nameoftool "GoFunctionSignature"` and it prints generated code to stdout.

Example code to be generated (WORK IN PROGRESS):

    struct exampleWrapper {
        f func(arg1, arg2, etc...) (out1, out2, etc..)
    }

    NewWrapper(f func(args...) (outs...)){
        wrapper = new exampleWrapper
        wrapper.f = f
        return wrapper
    }

    // Implements a funcWrapper interface, to be added to the labassistant library?
    wrapper.Run(arg1, arg2, etc...){
        outs = wrapper.f(arg1, arg2, etc..)
        return outs
    }


# GoMat
my hand-rolled library for matrix operations in Go.  Making this as an exercise for making usable libraries.  A side-goal of the project is to be able to make functions concurrent when it is possible

This is mainly an exercise in designing a library, so my current design process is to make the naive implementation, test the constraints of that, then try to optimize it based on the slower cases (typically large matrices).

As a result of this design process, this will in all likelihood not conform to the BLAS specification anytime soon, although I'd like to eventually get it there.

# :small_red_triangle: Debugo

*An optimized debug library for go*

Debugo is a dead simple debugging library that doesn't increase your programs size thanks to the compiler optimizations.

## Installation

```
go get github.com/negrel/debugo
```

## Hello world

**debugo.Assert** is a function that assert the given bool and panic if it is not true.

Take a look at the following example :

```go
package main

import (
	"github.com/negrel/debugo"
)

func main() {
	var hello string = "Hello world!"
	var bye string = "Bye Bye"

	// Obviously the following strings are not equal.
	fmt.Printf("%s == %s => %v", hello, bye, hello == bye)

	// Should panic if the given bool is false
	// and if the 'debugo' build tag is used.
	debugo.Assert(hello == bye, "error, the given string are not equal")
}
```

Let's run the program:

```
go run .
'Hello world!' == 'Bye Bye' is false
```

As you can see, no panic was triggered because the **debugo.Assert** function was not executed. The function was **removed** when the program was compiled.

Let's run the program with the debugo tag:
```
go run -tags debugo .
'Hello world!' == 'Bye Bye' is false
panic: error, the given string are not equal

goroutine 1 [running]:
github.com/negrel/debugo.Assert(...)
        /home/negrel/code/golang/src/github.com/negrel/beta/vendor/github.com/negrel/debugo/assert_d.go:10
main.main()
        /home/negrel/code/golang/src/github.com/negrel/beta/main.go:17 +0x146
exit status 2
```

Now the program panic because the compiler include the **debugo.Assert** function.

Let's take a look at the **Static Single Assignment**:
<table>
<tr>
<td>
<b>WITHOUT</b> debugo build tag
</td>
<td>
<b>WITH</b> debugo build tag
</td>
</tr>
<td>
<code><dl ><dt ></dt><dd># /home/negrel/code/golang/src/github.com/negrel/beta/main.go</dd><dt ></dt><dd >00000 <span >(9)</span> TEXT	"".main(SB), ABIInternal</dd><dt ></dt><dd >00001 <span >(9)</span> PCDATA	$0, $-2</dd><dt ></dt><dd >00002 <span >(9)</span> PCDATA	$1, $-2</dd><dt ></dt><dd >00003 <span >(9)</span> FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)</dd><dt ></dt><dd >00004 <span >(9)</span> FUNCDATA	$1, gclocals·cc6adf73a0f41ffa3266b3b03b2771b8(SB)</dd><dt ></dt><dd >00005 <span >(9)</span> FUNCDATA	$2, gclocals·3639c5e889acb2527c3873192ba4ec61(SB)</dd><dt ></dt><dd >00006 <span >(9)</span> FUNCDATA	$3, "".main.stkobj(SB)</dd><dt ><span >v50</span></dt><dd >00007 <span >(<b>+14</b>)</span> PCDATA	$0, $1</dd><dt ><span >v50</span></dt><dd >00008 <span >(<b>+14</b>)</span> PCDATA	$1, $0</dd><dt ><span >v50</span></dt><dd >00009 <span >(<b>+14</b>)</span> LEAQ	go.string."Hello world!"(SB), AX</dd><dt ><span >v60</span></dt><dd >00010 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v60</span></dt><dd >00011 <span >(<s>14</s>)</span> MOVQ	AX, (SP)</dd><dt ><span >v31</span></dt><dd >00012 <span >(<s>14</s>)</span> MOVQ	$12, 8(SP)</dd><dt ><span >v32</span></dt><dd >00013 <span >(<s>14</s>)</span> CALL	runtime.convTstring(SB)</dd><dt ><span >v34</span></dt><dd >00014 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v34</span></dt><dd >00015 <span >(<s>14</s>)</span> MOVQ	16(SP), AX</dd><dt ><span >v72</span></dt><dd >00016 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v72</span></dt><dd >00017 <span >(<s>14</s>)</span> PCDATA	$1, $1</dd><dt ><span >v72</span></dt><dd >00018 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_38-56(SP)</dd><dt ><span >v74</span></dt><dd >00019 <span >(<s>14</s>)</span> PCDATA	$0, $2</dd><dt ><span >v74</span></dt><dd >00020 <span >(<s>14</s>)</span> LEAQ	go.string."Bye Bye"(SB), CX</dd><dt ><span >v41</span></dt><dd >00021 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v41</span></dt><dd >00022 <span >(<s>14</s>)</span> MOVQ	CX, (SP)</dd><dt ><span >v38</span></dt><dd >00023 <span >(<s>14</s>)</span> MOVQ	$7, 8(SP)</dd><dt ><span >v39</span></dt><dd >00024 <span >(<s>14</s>)</span> CALL	runtime.convTstring(SB)</dd><dt ><span >v40</span></dt><dd >00025 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v40</span></dt><dd >00026 <span >(<s>14</s>)</span> MOVQ	16(SP), AX</dd><dt ><span >v36</span></dt><dd >00027 <span >(<s>14</s>)</span> PCDATA	$1, $2</dd><dt ><span >v36</span></dt><dd >00028 <span >(<s>14</s>)</span> XORPS	X0, X0</dd><dt ><span >v68</span></dt><dd >00029 <span >(<s>14</s>)</span> MOVUPS	X0, ""..autotmp_21-48(SP)</dd><dt ><span >v42</span></dt><dd >00030 <span >(<s>14</s>)</span> MOVUPS	X0, ""..autotmp_21-32(SP)</dd><dt ><span >v53</span></dt><dd >00031 <span >(<s>14</s>)</span> MOVUPS	X0, ""..autotmp_21-16(SP)</dd><dt ><span >v82</span></dt><dd >00032 <span >(<s>14</s>)</span> PCDATA	$0, $3</dd><dt ><span >v82</span></dt><dd >00033 <span >(<s>14</s>)</span> LEAQ	type.string(SB), CX</dd><dt ><span >v111</span></dt><dd >00034 <span >(<s>14</s>)</span> MOVQ	CX, ""..autotmp_21-48(SP)</dd><dt ><span >v95</span></dt><dd >00035 <span >(<s>14</s>)</span> PCDATA	$0, $4</dd><dt ><span >v95</span></dt><dd >00036 <span >(<s>14</s>)</span> PCDATA	$1, $3</dd><dt ><span >v95</span></dt><dd >00037 <span >(<s>14</s>)</span> MOVQ	""..autotmp_38-56(SP), DX</dd><dt ><span >v59</span></dt><dd >00038 <span >(<s>14</s>)</span> PCDATA	$0, $3</dd><dt ><span >v59</span></dt><dd >00039 <span >(<s>14</s>)</span> MOVQ	DX, ""..autotmp_21-40(SP)</dd><dt ><span >v108</span></dt><dd >00040 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v108</span></dt><dd >00041 <span >(<s>14</s>)</span> MOVQ	CX, ""..autotmp_21-32(SP)</dd><dt ><span >v63</span></dt><dd >00042 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v63</span></dt><dd >00043 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_21-24(SP)</dd><dt ><span >v100</span></dt><dd >00044 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v100</span></dt><dd >00045 <span >(<s>14</s>)</span> LEAQ	type.bool(SB), AX</dd><dt ><span >v106</span></dt><dd >00046 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v106</span></dt><dd >00047 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_21-16(SP)</dd><dt ><span >v104</span></dt><dd >00048 <span >(<s>14</s>)</span> XORL	AX, AX</dd><dt ><span >v47</span></dt><dd >00049 <span >(<s>14</s>)</span> MOVBLZX	AX, AX</dd><dt ><span >v107</span></dt><dd >00050 <span >(<s>14</s>)</span> PCDATA	$0, $2</dd><dt ><span >v107</span></dt><dd >00051 <span >(<s>14</s>)</span> LEAQ	runtime.staticbytes(SB), CX</dd><dt ><span >v48</span></dt><dd >00052 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v48</span></dt><dd >00053 <span >(<s>14</s>)</span> ADDQ	CX, AX</dd><dt ><span >v67</span></dt><dd >00054 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v67</span></dt><dd >00055 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_21-8(SP)</dd><dt ><span >v77</span></dt><dd >00056 <span>(?)</span> NOP</dd><dt ></dt><dd ># $GOROOT/src/fmt/print.go</dd><dt ><span >v81</span></dt><dd >00057 <span >(<b>+213</b>)</span> PCDATA	$0, $1</dd><dt ><span >v81</span></dt><dd >00058 <span >(<b>+213</b>)</span> MOVQ	os.Stdout(SB), AX</dd><dt ><span >v110</span></dt><dd >00059 <span >(<s>213</s>)</span> PCDATA	$0, $3</dd><dt ><span >v110</span></dt><dd >00060 <span >(<s>213</s>)</span> LEAQ	go.itab.*os.File,io.Writer(SB), CX</dd><dt ><span >v102</span></dt><dd >00061 <span >(<s>213</s>)</span> PCDATA	$0, $1</dd><dt ><span >v102</span></dt><dd >00062 <span >(<s>213</s>)</span> MOVQ	CX, (SP)</dd><dt ><span >v83</span></dt><dd >00063 <span >(<s>213</s>)</span> PCDATA	$0, $0</dd><dt ><span >v83</span></dt><dd >00064 <span >(<s>213</s>)</span> MOVQ	AX, 8(SP)</dd><dt ><span >v115</span></dt><dd >00065 <span >(<s>213</s>)</span> PCDATA	$0, $1</dd><dt ><span >v115</span></dt><dd >00066 <span >(<s>213</s>)</span> LEAQ	go.string."'%s' == '%v' is %v\n"(SB), AX</dd><dt ><span >v98</span></dt><dd >00067 <span >(<s>213</s>)</span> PCDATA	$0, $0</dd><dt ><span >v98</span></dt><dd >00068 <span >(<s>213</s>)</span> MOVQ	AX, 16(SP)</dd><dt ><span >v86</span></dt><dd >00069 <span >(<s>213</s>)</span> MOVQ	$19, 24(SP)</dd><dt ><span >v55</span></dt><dd >00070 <span >(<s>213</s>)</span> PCDATA	$0, $1</dd><dt ><span >v55</span></dt><dd >00071 <span >(<s>213</s>)</span> PCDATA	$1, $0</dd><dt ><span >v55</span></dt><dd >00072 <span >(<s>213</s>)</span> LEAQ	""..autotmp_21-48(SP), AX</dd><dt ><span >v91</span></dt><dd >00073 <span >(<s>213</s>)</span> PCDATA	$0, $0</dd><dt ><span >v91</span></dt><dd >00074 <span >(<s>213</s>)</span> MOVQ	AX, 32(SP)</dd><dt ><span >v73</span></dt><dd >00075 <span >(<s>213</s>)</span> MOVQ	$3, 40(SP)</dd><dt ><span >v88</span></dt><dd >00076 <span >(<s>213</s>)</span> MOVQ	$3, 48(SP)</dd><dt ><span >v89</span></dt><dd >00077 <span >(<s>213</s>)</span> CALL	fmt.Fprintf(SB)</dd><dt ></dt><dd ># /home/negrel/code/golang/src/github.com/negrel/beta/main.go</dd><dt ><b><span >b9</span></b></dt><dd >00078 <span >(<b>+17</b>)</span> RET</dd><dt ></dt><dd >00079 <span>(?)</span> END</dd></dl></code>
</td>
<td>
<code><dl ><dt ></dt><dd ># /home/negrel/code/golang/src/github.com/negrel/beta/main.go</dd><dt ></dt><dd >00000 <span >(9)</span> TEXT	"".main(SB), ABIInternal</dd><dt ></dt><dd >00001 <span >(9)</span> PCDATA	$0, $-2</dd><dt ></dt><dd >00002 <span >(9)</span> PCDATA	$1, $-2</dd><dt ></dt><dd >00003 <span >(9)</span> FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)</dd><dt ></dt><dd >00004 <span >(9)</span> FUNCDATA	$1, gclocals·cc6adf73a0f41ffa3266b3b03b2771b8(SB)</dd><dt ></dt><dd >00005 <span >(9)</span> FUNCDATA	$2, gclocals·1ecaf0bbee428656b8c0e1f08a3a1364(SB)</dd><dt ></dt><dd >00006 <span >(9)</span> FUNCDATA	$3, "".main.stkobj(SB)</dd><dt ><span >v68</span></dt><dd >00007 <span >(<b>+14</b>)</span> PCDATA	$0, $1</dd><dt ><span >v68</span></dt><dd >00008 <span >(<b>+14</b>)</span> PCDATA	$1, $0</dd><dt ><span >v68</span></dt><dd >00009 <span >(<b>+14</b>)</span> LEAQ	go.string."Hello world!"(SB), AX</dd><dt ><span >v60</span></dt><dd >00010 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v60</span></dt><dd >00011 <span >(<s>14</s>)</span> MOVQ	AX, (SP)</dd><dt ><span >v31</span></dt><dd >00012 <span >(<s>14</s>)</span> MOVQ	$12, 8(SP)</dd><dt ><span >v32</span></dt><dd >00013 <span >(<s>14</s>)</span> CALL	runtime.convTstring(SB)</dd><dt ><span >v34</span></dt><dd >00014 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v34</span></dt><dd >00015 <span >(<s>14</s>)</span> MOVQ	16(SP), AX</dd><dt ><span >v45</span></dt><dd >00016 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v45</span></dt><dd >00017 <span >(<s>14</s>)</span> PCDATA	$1, $1</dd><dt ><span >v45</span></dt><dd >00018 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_44-56(SP)</dd><dt ><span >v82</span></dt><dd >00019 <span >(<s>14</s>)</span> PCDATA	$0, $2</dd><dt ><span >v82</span></dt><dd >00020 <span >(<s>14</s>)</span> LEAQ	go.string."Bye Bye"(SB), CX</dd><dt ><span >v123</span></dt><dd >00021 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v123</span></dt><dd >00022 <span >(<s>14</s>)</span> MOVQ	CX, (SP)</dd><dt ><span >v38</span></dt><dd >00023 <span >(<s>14</s>)</span> MOVQ	$7, 8(SP)</dd><dt ><span >v39</span></dt><dd >00024 <span >(<s>14</s>)</span> CALL	runtime.convTstring(SB)</dd><dt ><span >v40</span></dt><dd >00025 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v40</span></dt><dd >00026 <span >(<s>14</s>)</span> MOVQ	16(SP), AX</dd><dt ><span >v42</span></dt><dd >00027 <span >(<s>14</s>)</span> PCDATA	$1, $2</dd><dt ><span >v42</span></dt><dd >00028 <span >(<s>14</s>)</span> XORPS	X0, X0</dd><dt ><span >v49</span></dt><dd >00029 <span >(<s>14</s>)</span> MOVUPS	X0, ""..autotmp_24-48(SP)</dd><dt ><span >v74</span></dt><dd >00030 <span >(<s>14</s>)</span> MOVUPS	X0, ""..autotmp_24-32(SP)</dd><dt ><span >v53</span></dt><dd >00031 <span >(<s>14</s>)</span> MOVUPS	X0, ""..autotmp_24-16(SP)</dd><dt ><span >v124</span></dt><dd >00032 <span >(<s>14</s>)</span> PCDATA	$0, $3</dd><dt ><span >v124</span></dt><dd >00033 <span >(<s>14</s>)</span> LEAQ	type.string(SB), CX</dd><dt ><span >v115</span></dt><dd >00034 <span >(<s>14</s>)</span> MOVQ	CX, ""..autotmp_24-48(SP)</dd><dt ><span >v73</span></dt><dd >00035 <span >(<s>14</s>)</span> PCDATA	$0, $4</dd><dt ><span >v73</span></dt><dd >00036 <span >(<s>14</s>)</span> PCDATA	$1, $3</dd><dt ><span >v73</span></dt><dd >00037 <span >(<s>14</s>)</span> MOVQ	""..autotmp_44-56(SP), DX</dd><dt ><span >v59</span></dt><dd >00038 <span >(<s>14</s>)</span> PCDATA	$0, $3</dd><dt ><span >v59</span></dt><dd >00039 <span >(<s>14</s>)</span> MOVQ	DX, ""..autotmp_24-40(SP)</dd><dt ><span >v112</span></dt><dd >00040 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v112</span></dt><dd >00041 <span >(<s>14</s>)</span> MOVQ	CX, ""..autotmp_24-32(SP)</dd><dt ><span >v63</span></dt><dd >00042 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v63</span></dt><dd >00043 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_24-24(SP)</dd><dt ><span >v98</span></dt><dd >00044 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v98</span></dt><dd >00045 <span >(<s>14</s>)</span> LEAQ	type.bool(SB), AX</dd><dt ><span >v110</span></dt><dd >00046 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v110</span></dt><dd >00047 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_24-16(SP)</dd><dt ><span >v102</span></dt><dd >00048 <span >(<s>14</s>)</span> XORL	AX, AX</dd><dt ><span >v47</span></dt><dd >00049 <span >(<s>14</s>)</span> MOVBLZX	AX, AX</dd><dt ><span >v106</span></dt><dd >00050 <span >(<s>14</s>)</span> PCDATA	$0, $5</dd><dt ><span >v106</span></dt><dd >00051 <span >(<s>14</s>)</span> LEAQ	runtime.staticbytes(SB), DX</dd><dt ><span >v48</span></dt><dd >00052 <span >(<s>14</s>)</span> PCDATA	$0, $1</dd><dt ><span >v48</span></dt><dd >00053 <span >(<s>14</s>)</span> ADDQ	DX, AX</dd><dt ><span >v67</span></dt><dd >00054 <span >(<s>14</s>)</span> PCDATA	$0, $0</dd><dt ><span >v67</span></dt><dd >00055 <span >(<s>14</s>)</span> MOVQ	AX, ""..autotmp_24-8(SP)</dd><dt ><span >v77</span></dt><dd >00056 <span>(?)</span> NOP</dd><dt ></dt><dd ># $GOROOT/src/fmt/print.go</dd><dt ><span >v81</span></dt><dd >00057 <span >(<b>+213</b>)</span> PCDATA	$0, $1</dd><dt ><span >v81</span></dt><dd >00058 <span >(<b>+213</b>)</span> MOVQ	os.Stdout(SB), AX</dd><dt ><span >v108</span></dt><dd >00059 <span >(<s>213</s>)</span> PCDATA	$0, $6</dd><dt ><span >v108</span></dt><dd >00060 <span >(<s>213</s>)</span> LEAQ	go.itab.*os.File,io.Writer(SB), DX</dd><dt ><span >v107</span></dt><dd >00061 <span >(<s>213</s>)</span> PCDATA	$0, $1</dd><dt ><span >v107</span></dt><dd >00062 <span >(<s>213</s>)</span> MOVQ	DX, (SP)</dd><dt ><span >v83</span></dt><dd >00063 <span >(<s>213</s>)</span> PCDATA	$0, $0</dd><dt ><span >v83</span></dt><dd >00064 <span >(<s>213</s>)</span> MOVQ	AX, 8(SP)</dd><dt ><span >v111</span></dt><dd >00065 <span >(<s>213</s>)</span> PCDATA	$0, $1</dd><dt ><span >v111</span></dt><dd >00066 <span >(<s>213</s>)</span> LEAQ	go.string."'%v' == '%s' is %v\n"(SB), AX</dd><dt ><span >v104</span></dt><dd >00067 <span >(<s>213</s>)</span> PCDATA	$0, $0</dd><dt ><span >v104</span></dt><dd >00068 <span >(<s>213</s>)</span> MOVQ	AX, 16(SP)</dd><dt ><span >v86</span></dt><dd >00069 <span >(<s>213</s>)</span> MOVQ	$19, 24(SP)</dd><dt ><span >v113</span></dt><dd >00070 <span >(<s>213</s>)</span> PCDATA	$0, $1</dd><dt ><span >v113</span></dt><dd >00071 <span >(<s>213</s>)</span> PCDATA	$1, $0</dd><dt ><span >v113</span></dt><dd >00072 <span >(<s>213</s>)</span> LEAQ	""..autotmp_24-48(SP), AX</dd><dt ><span >v95</span></dt><dd >00073 <span >(<s>213</s>)</span> PCDATA	$0, $0</dd><dt ><span >v95</span></dt><dd >00074 <span >(<s>213</s>)</span> MOVQ	AX, 32(SP)</dd><dt ><span >v100</span></dt><dd >00075 <span >(<s>213</s>)</span> MOVQ	$3, 40(SP)</dd><dt ><span >v88</span></dt><dd >00076 <span >(<s>213</s>)</span> MOVQ	$3, 48(SP)</dd><dt ><span >v89</span></dt><dd >00077 <span >(<s>213</s>)</span> CALL	fmt.Fprintf(SB)</dd><dt ></dt><dd ># /home/negrel/code/golang/src/github.com/negrel/beta/main.go</dd><dt ><span >v116</span></dt><dd >00078 <span >(<b>+17</b>)</span> XCHGL	AX, AX</dd><dt ></dt><dd ># /home/negrel/code/golang/src/github.com/negrel/beta/vendor/github.com/negrel/debugo/assert_d.go</dd><dt ><span >v41</span></dt><dd >00079 <span >(<s>10</s>)</span> PCDATA	$0, $1</dd><dt ><span >v41</span></dt><dd >00080 <span >(<b>+10</b>)</span> LEAQ	go.string."error, the given string are not equal"(SB), AX</dd><dt ><span >v72</span></dt><dd >00081 <span >(<s>10</s>)</span> PCDATA	$0, $0</dd><dt ><span >v72</span></dt><dd >00082 <span >(<s>10</s>)</span> MOVQ	AX, (SP)</dd><dt ><span >v119</span></dt><dd >00083 <span >(<s>10</s>)</span> MOVQ	$37, 8(SP)</dd><dt ><span >v120</span></dt><dd >00084 <span >(<b>+10</b>)</span> CALL	runtime.convTstring(SB)</dd><dt ><span >v121</span></dt><dd >00085 <span >(<s>10</s>)</span> PCDATA	$0, $1</dd><dt ><span >v121</span></dt><dd >00086 <span >(<s>10</s>)</span> MOVQ	16(SP), AX</dd><dt ><span >v55</span></dt><dd >00087 <span >(<s>10</s>)</span> PCDATA	$0, $3</dd><dt ><span >v55</span></dt><dd >00088 <span >(<s>10</s>)</span> LEAQ	type.string(SB), CX</dd><dt ><span >v69</span></dt><dd >00089 <span >(<s>10</s>)</span> PCDATA	$0, $1</dd><dt ><span >v69</span></dt><dd >00090 <span >(<s>10</s>)</span> MOVQ	CX, (SP)</dd><dt ><span >v125</span></dt><dd >00091 <span >(<s>10</s>)</span> PCDATA	$0, $0</dd><dt ><span >v125</span></dt><dd >00092 <span >(<s>10</s>)</span> MOVQ	AX, 8(SP)</dd><dt ><span >v126</span></dt><dd >00093 <span >(<s>10</s>)</span> CALL	runtime.gopanic(SB)</dd><dt ></dt><dd >00094 <span >(10)</span> XCHGL	AX, AX</dd><dt ></dt><dd >00095 <span>(?)</span> END</dd></dl></code>
</td>
</tr>

</table>

### Contributing
If you want to contribute to Ginger to add a feature or improve the code contact me at [negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/debugo/issues) or make a [pull request](https://github.com/negrel/debugo/pulls).

## :stars: Show your support
Please give a :star: if this project helped you!

#### :scroll: License
MIT © Alexandre Negrel

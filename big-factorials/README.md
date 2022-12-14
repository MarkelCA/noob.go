# Big factorials

This is a test to check the performance difference between go and an interpreted language like Python on big calculations. These will consist on timing a couple of factorial functions with big numbers

These function execution measures just the calculation time, not the printing of the number. This test are done with the factorial of two million (2.000.000).

## Test results

### Go

Command:
```bash
time go run fac.go
```
Output:
```
Executed in    5,20 secs   fish           external
   usr time    5,22 secs  648,00 micros    5,22 secs
   sys time    0,09 secs    0,00 micros    0,09 secs
```

### Python
Command:
```bash
time python3 fac.py 
```
Output:
```
Executed in   18,56 secs   fish           external
   usr time   18,52 secs  614,00 micros   18,52 secs
   sys time    0,04 secs  188,00 micros    0,04 secs
```

### PyPy
Done with a Just In Time compiled implementation of Python. Took nearly half of the time
Command:
```bash
time pypy fac.py
```
Output:
```
Executed in    9,33 secs   fish           external
   usr time    9,28 secs    0,00 micros    9,28 secs
   sys time    0,02 secs  804,00 micros    0,02 secs
```

### Java
Command:
```bash
time java fac.java
```
Output:
```
Executed in   20,26 mins   fish           external 
   usr time  1258,22 secs  401,00 micros  1258,22 secs 
   sys time   13,83 secs  120,00 micros   13,83 secs 

```
I'm not sure why it took so long comparing to Python

## CPU
Just for the record. My CPU specifications is the following.
```
Architecture:                    x86_64
CPU op-mode(s):                  32-bit, 64-bit
Byte Order:                      Little Endian
Address sizes:                   39 bits physical, 48 bits virtual
CPU(s):                          8
On-line CPU(s) list:             0-7
Thread(s) per core:              2
Core(s) per socket:              4
Socket(s):                       1
NUMA node(s):                    1
Vendor ID:                       GenuineIntel
CPU family:                      6
Model:                           140
Model name:                      11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Stepping:                        1
CPU MHz:                         2400.000
CPU max MHz:                     4200,0000
CPU min MHz:                     400,0000
BogoMIPS:                        4838.40
Virtualization:                  VT-x
L1d cache:                       192 KiB
L1i cache:                       128 KiB
L2 cache:                        5 MiB
L3 cache:                        8 MiB
NUMA node0 CPU(s):               0-7
Vulnerability Itlb multihit:     Not affected
Vulnerability L1tf:              Not affected
Vulnerability Mds:               Not affected
Vulnerability Meltdown:          Not affected
Vulnerability Spec store bypass: Mitigation; Speculative Store Bypass disabled via prctl and seccomp
Vulnerability Spectre v1:        Mitigation; usercopy/swapgs barriers and __user pointer sanitization
Vulnerability Spectre v2:        Mitigation; Enhanced IBRS, IBPB conditional, RSB filling
Vulnerability Srbds:             Not affected
Vulnerability Tsx async abort:   Not affected
Flags:                           fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon p
                                 ebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf tsc_known_freq pni pclmulqdq dtes64 monitor ds_cpl vmx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic
                                 movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault epb cat_l2 invpcid_single cdp_l2 ssbd ibrs ibpb stibp ibrs_enhanced tpr_shadow vnmi
                                 flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb intel_pt avx512cd sha_ni avx512b
                                 w avx512vl xsaveopt xsavec xgetbv1 xsaves split_lock_detect dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp hwp_pkg_req avx512vbmi umip pku ospke avx512_vbmi2 gfni va
                                 es vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq rdpid movdiri movdir64b fsrm avx512_vp2intersect md_clear flush_l1d arch_capabilities
```

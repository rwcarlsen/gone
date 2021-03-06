package endf

var mtname = map[int]string{
	1:   "(n,total) neutron total",
	2:   "(z,z0) elastic scattering",
	3:   "(z,nonelas) nonelastic neutron",
	4:   "(z,n) one neutron in exit channel",
	5:   "(z,anything) miscellaneous",
	10:  "(z,contin) total continuum reaction",
	11:  "(z,2nd) production of 2n and d",
	16:  "(z,2n) production of 2n",
	17:  "(z,3n) production of 3n",
	18:  "(z,fiss) particle-induced fission",
	19:  "(z,f) first-chance fission",
	20:  "(z,nf) second chance fission",
	21:  "(z,2nf) third-chance fission",
	22:  "(z,na) production of n and alpha",
	23:  "(z,n3a) production of n and 3 alphas",
	24:  "(z,2na) production of 2n and alpha",
	25:  "(z,3na) production of 3n and alpha",
	27:  "(n,abs) absorption",
	28:  "(z,np) production of n and p",
	29:  "(z,n2a) production of n and 2 alphas",
	30:  "(z,2n2a) production of 2n and 2 alphas",
	32:  "(z,nd) production of n and d",
	33:  "(z,nt) production of n and t",
	34:  "(z,n3he) production of n and he-3",
	35:  "(z,nd2a) production of n, d, and alpha",
	36:  "(z,nt2a) production of n, t, and 2 alphas",
	37:  "(z,4n) production of 4n",
	38:  "(z,3nf) fourth-chance fission",
	41:  "(z,2np) production of 2n and p",
	42:  "(z,3np) production of 3n and p",
	44:  "(z,n2p) production of n and 2p",
	45:  "(z,npa) production of n, p, and alpha",
	50:  "(z,n0) production of n, ground state",
	51:  "(z,n1) production of n, 1st excited state",
	52:  "(z,n2) production of n, 2nd excited state",
	53:  "(z,n3) production of n, 3rd excited state",
	54:  "(z,n4) production of n, 4th excited state",
	55:  "(z,n5) production of n, 5th excited state",
	56:  "(z,n6) production of n, 6th excited state",
	57:  "(z,n7) production of n, 7th excited state",
	58:  "(z,n8) production of n, 8th excited state",
	59:  "(z,n9) production of n, 9th excited state",
	60:  "(z,n10) production of n, 10th excited state",
	61:  "(z,n11) production of n, 11th excited state",
	62:  "(z,n12) production of n, 12th excited state",
	63:  "(z,n13) production of n, 13th excited state",
	64:  "(z,n14) production of n, 14th excited state",
	65:  "(z,n15) production of n, 15th excited state",
	66:  "(z,n16) production of n, 16th excited state",
	67:  "(z,n17) production of n, 17th excited state",
	68:  "(z,n18) production of n, 18th excited state",
	69:  "(z,n19) production of n, 19th excited state",
	70:  "(z,n20) production of n, 20th excited state",
	71:  "(z,n21) production of n, 21st excited state",
	72:  "(z,n22) production of n, 22nd excited state",
	73:  "(z,n23) production of n, 23rd excited state",
	74:  "(z,n24) production of n, 24th excited state",
	75:  "(z,n25) production of n, 25th excited state",
	76:  "(z,n26) production of n, 26th excited state",
	77:  "(z,n27) production of n, 27th excited state",
	78:  "(z,n28) production of n, 28th excited state",
	79:  "(z,n29) production of n, 29th excited state",
	80:  "(z,n30) production of n, 30th excited state",
	81:  "(z,n31) production of n, 31st excited state",
	82:  "(z,n32) production of n, 32nd excited state",
	83:  "(z,n33) production of n, 33rd excited state",
	84:  "(z,n34) production of n, 34th excited state",
	85:  "(z,n35) production of n, 35th excited state",
	86:  "(z,n36) production of n, 36th excited state",
	87:  "(z,n37) production of n, 37th excited state",
	88:  "(z,n38) production of n, 38th excited state",
	89:  "(z,n39) production of n, 39th excited state",
	90:  "(z,n40) production of n, 40th excited state",
	91:  "(z,nc) production of n in continuum",
	101: "(n,disap) neutron disappeareance",
	102: "(z,gamma) radiative capture",
	103: "(z,p) production of p",
	104: "(z,d) production of d",
	105: "(z,t) production of t",
	106: "(z,3he) production of he-3",
	107: "(z,a) production of alpha",
	108: "(z,2a) production of 2 alphas",
	109: "(z,3a) production of 3 alphas",
	111: "(z,2p) production of 2p",
	112: "(z,pa) production of p and alpha",
	113: "(z,t2a) production of t and 2 alphas",
	114: "(z,d2a) production of d and 2 alphas",
	115: "(z,pd) production of p and d",
	116: "(z,pt) production of p and t",
	117: "(z,da) production of d and a",
	151: "resonance parameters",
	201: "(z,xn) total neutron production",
	202: "(z,xgamma) total gamma production",
	203: "(z,xp) total proton production",
	204: "(z,xd) total deuteron production",
	205: "(z,xt) total triton production",
	206: "(z,x3he) total he-3 production",
	207: "(z,xa) total alpha production",
	208: "(z,xpi+) total pi+ meson production",
	209: "(z,xpi0) total pi0 meson production",
	210: "(z,xpi-) total pi- meson production",
	211: "(z,xmu+) total anti-muon production",
	212: "(z,xmu-) total muon production",
	213: "(z,xk+) total positive kaon production",
	214: "(z,xk0long) total long-lived neutral kaon production",
	215: "(z,xk0short) total short-lived neutral kaon production",
	216: "(z,xk-) total negative kaon production",
	217: "(z,xp-) total anti-proton production",
	218: "(z,xn-) total anti-neutron production",
	251: "average cosine of scattering angle",
	252: "average logarithmic energy decrement",
	253: "average xi^2/(2*xi)",
	451: "desciptive data",
	452: "total neutrons per fission",
	454: "independent fission product yield",
	455: "delayed neutron data",
	456: "prompt neutrons per fission",
	457: "radioactive decay data",
	458: "energy release due to fission",
	459: "cumulative fission product yield",
	460: "delayed photon data",
	500: "total charged-particle stopping power",
	501: "total photon interaction",
	502: "photon coherent scattering",
	504: "photon incoherent scattering",
	505: "imaginary scattering factor",
	506: "real scattering factor",
	515: "pair production, electron field",
	516: "total pair production",
	517: "pair production, nuclear field",
	522: "photoelectric absorption",
	523: "photo-excitation cross section",
	526: "electro-atomic scattering",
	527: "electro-atomic bremsstrahlung",
	528: "electro-atomic excitation cross section",
	533: "atomic relaxation data",
	534: "k (1s1/2) subshell",
	535: "l1 (2s1/2) subshell",
	536: "l2 (2p1/2) subshell",
	537: "l3 (2p3/2) subshell",
	538: "m1 (3s1/2) subshell",
	539: "m2 (3p1/2) subshell",
	540: "m3 (3p3/2) subshell",
	541: "m4 (3d1/2) subshell",
	542: "m5 (3d1/2) subshell",
	543: "n1 (4s1/2) subshell",
	544: "n2 (4p1/2) subshell",
	545: "n3 (4p3/2) subshell",
	546: "n4 (4d3/2) subshell",
	547: "n5 (4d5/2) subshell",
	548: "n6 (4f5/2) subshell",
	549: "n7 (4f7/2) subshell",
	550: "o1 (5s1/2) subshell",
	551: "o2 (5p1/2) subshell",
	552: "o3 (5p3/2) subshell",
	553: "o4 (5d3/2) subshell",
	554: "o5 (5d5/2) subshell",
	555: "o6 (5f5/2) subshell",
	556: "o7 (5f7/2) subshell",
	557: "o8 (5g7/2) subshell",
	558: "o9 (5g9/2) subshell",
	559: "p1 (6s1/2) subshell",
	560: "p2 (6p1/2) subshell",
	561: "p3 (6p3/2) subshell",
	562: "p4 (6d3/2) subshell",
	563: "p5 (6d5/2) subshell",
	564: "p6 (6f5/2) subshell",
	565: "p7 (6f7/2) subshell",
	566: "p8 (6g7/2) subshell",
	567: "p9 (6g9/2) subshell",
	568: "p10 (6h9/2) subshell",
	569: "p11 (6h11/2) subshell",
	570: "q1 (7s1/2) subshell",
	571: "q2 (7p1/2) subshell",
	572: "q3 (7p3/2) subshell",
}

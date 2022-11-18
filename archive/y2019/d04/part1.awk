#GNU Awk 5.0.1, API: 2.0 (GNU MPFR 4.0.2, GNU MP 6.1.2)
#Copyright (C) 1989, 1991-2019 Free Software Foundation.

BEGIN {
	FS = ""
};

$1<=$2 &&
$2<=$3 &&
$3<=$4 &&
$4<=$5 &&
$5<=$6 &&
($1==$2 ||
$2==$3 ||
$3==$4 ||
$4==$5 ||
$5==$6) {
	print $0
}

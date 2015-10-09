package treasures

const (
	md5js = `/**
*
*  MD5 (Message-Digest Algorithm)
*  http://www.webtoolkit.info/
*
**/

var MD5 = function (string) {
 
	function RotateLeft(lValue, iShiftBits) {
		return (lValue<<iShiftBits) | (lValue>>>(32-iShiftBits));
	}
 
	function AddUnsigned(lX,lY) {
		var lX4,lY4,lX8,lY8,lResult;
		lX8 = (lX & 0x80000000);
		lY8 = (lY & 0x80000000);
		lX4 = (lX & 0x40000000);
		lY4 = (lY & 0x40000000);
		lResult = (lX & 0x3FFFFFFF)+(lY & 0x3FFFFFFF);
		if (lX4 & lY4) {
			return (lResult ^ 0x80000000 ^ lX8 ^ lY8);
		}
		if (lX4 | lY4) {
			if (lResult & 0x40000000) {
				return (lResult ^ 0xC0000000 ^ lX8 ^ lY8);
			} else {
				return (lResult ^ 0x40000000 ^ lX8 ^ lY8);
			}
		} else {
			return (lResult ^ lX8 ^ lY8);
		}
 	}
 
 	function F(x,y,z) { return (x & y) | ((~x) & z); }
 	function G(x,y,z) { return (x & z) | (y & (~z)); }
 	function H(x,y,z) { return (x ^ y ^ z); }
	function I(x,y,z) { return (y ^ (x | (~z))); }
 
	function FF(a,b,c,d,x,s,ac) {
		a = AddUnsigned(a, AddUnsigned(AddUnsigned(F(b, c, d), x), ac));
		return AddUnsigned(RotateLeft(a, s), b);
	};
 
	function GG(a,b,c,d,x,s,ac) {
		a = AddUnsigned(a, AddUnsigned(AddUnsigned(G(b, c, d), x), ac));
		return AddUnsigned(RotateLeft(a, s), b);
	};
 
	function HH(a,b,c,d,x,s,ac) {
		a = AddUnsigned(a, AddUnsigned(AddUnsigned(H(b, c, d), x), ac));
		return AddUnsigned(RotateLeft(a, s), b);
	};
 
	function II(a,b,c,d,x,s,ac) {
		a = AddUnsigned(a, AddUnsigned(AddUnsigned(I(b, c, d), x), ac));
		return AddUnsigned(RotateLeft(a, s), b);
	};
 
	function ConvertToWordArray(string) {
		var lWordCount;
		var lMessageLength = string.length;
		var lNumberOfWords_temp1=lMessageLength + 8;
		var lNumberOfWords_temp2=(lNumberOfWords_temp1-(lNumberOfWords_temp1 % 64))/64;
		var lNumberOfWords = (lNumberOfWords_temp2+1)*16;
		var lWordArray=Array(lNumberOfWords-1);
		var lBytePosition = 0;
		var lByteCount = 0;
		while ( lByteCount < lMessageLength ) {
			lWordCount = (lByteCount-(lByteCount % 4))/4;
			lBytePosition = (lByteCount % 4)*8;
			lWordArray[lWordCount] = (lWordArray[lWordCount] | (string.charCodeAt(lByteCount)<<lBytePosition));
			lByteCount++;
		}
		lWordCount = (lByteCount-(lByteCount % 4))/4;
		lBytePosition = (lByteCount % 4)*8;
		lWordArray[lWordCount] = lWordArray[lWordCount] | (0x80<<lBytePosition);
		lWordArray[lNumberOfWords-2] = lMessageLength<<3;
		lWordArray[lNumberOfWords-1] = lMessageLength>>>29;
		return lWordArray;
	};
 
	function WordToHex(lValue) {
		var WordToHexValue="",WordToHexValue_temp="",lByte,lCount;
		for (lCount = 0;lCount<=3;lCount++) {
			lByte = (lValue>>>(lCount*8)) & 255;
			WordToHexValue_temp = "0" + lByte.toString(16);
			WordToHexValue = WordToHexValue + WordToHexValue_temp.substr(WordToHexValue_temp.length-2,2);
		}
		return WordToHexValue;
	};
 
	function Utf8Encode(string) {
		string = string.replace(/\r\n/g,"\n");
		var utftext = "";
 
		for (var n = 0; n < string.length; n++) {
 
			var c = string.charCodeAt(n);
 
			if (c < 128) {
				utftext += String.fromCharCode(c);
			}
			else if((c > 127) && (c < 2048)) {
				utftext += String.fromCharCode((c >> 6) | 192);
				utftext += String.fromCharCode((c & 63) | 128);
			}
			else {
				utftext += String.fromCharCode((c >> 12) | 224);
				utftext += String.fromCharCode(((c >> 6) & 63) | 128);
				utftext += String.fromCharCode((c & 63) | 128);
			}
 
		}
 
		return utftext;
	};
 
	var x=Array();
	var k,AA,BB,CC,DD,a,b,c,d;
	var S11=7, S12=12, S13=17, S14=22;
	var S21=5, S22=9 , S23=14, S24=20;
	var S31=4, S32=11, S33=16, S34=23;
	var S41=6, S42=10, S43=15, S44=21;
 
	string = Utf8Encode(string);
 
	x = ConvertToWordArray(string);
 
	a = 0x67452301; b = 0xEFCDAB89; c = 0x98BADCFE; d = 0x10325476;
 
	for (k=0;k<x.length;k+=16) {
		AA=a; BB=b; CC=c; DD=d;
		a=FF(a,b,c,d,x[k+0], S11,0xD76AA478);
		d=FF(d,a,b,c,x[k+1], S12,0xE8C7B756);
		c=FF(c,d,a,b,x[k+2], S13,0x242070DB);
		b=FF(b,c,d,a,x[k+3], S14,0xC1BDCEEE);
		a=FF(a,b,c,d,x[k+4], S11,0xF57C0FAF);
		d=FF(d,a,b,c,x[k+5], S12,0x4787C62A);
		c=FF(c,d,a,b,x[k+6], S13,0xA8304613);
		b=FF(b,c,d,a,x[k+7], S14,0xFD469501);
		a=FF(a,b,c,d,x[k+8], S11,0x698098D8);
		d=FF(d,a,b,c,x[k+9], S12,0x8B44F7AF);
		c=FF(c,d,a,b,x[k+10],S13,0xFFFF5BB1);
		b=FF(b,c,d,a,x[k+11],S14,0x895CD7BE);
		a=FF(a,b,c,d,x[k+12],S11,0x6B901122);
		d=FF(d,a,b,c,x[k+13],S12,0xFD987193);
		c=FF(c,d,a,b,x[k+14],S13,0xA679438E);
		b=FF(b,c,d,a,x[k+15],S14,0x49B40821);
		a=GG(a,b,c,d,x[k+1], S21,0xF61E2562);
		d=GG(d,a,b,c,x[k+6], S22,0xC040B340);
		c=GG(c,d,a,b,x[k+11],S23,0x265E5A51);
		b=GG(b,c,d,a,x[k+0], S24,0xE9B6C7AA);
		a=GG(a,b,c,d,x[k+5], S21,0xD62F105D);
		d=GG(d,a,b,c,x[k+10],S22,0x2441453);
		c=GG(c,d,a,b,x[k+15],S23,0xD8A1E681);
		b=GG(b,c,d,a,x[k+4], S24,0xE7D3FBC8);
		a=GG(a,b,c,d,x[k+9], S21,0x21E1CDE6);
		d=GG(d,a,b,c,x[k+14],S22,0xC33707D6);
		c=GG(c,d,a,b,x[k+3], S23,0xF4D50D87);
		b=GG(b,c,d,a,x[k+8], S24,0x455A14ED);
		a=GG(a,b,c,d,x[k+13],S21,0xA9E3E905);
		d=GG(d,a,b,c,x[k+2], S22,0xFCEFA3F8);
		c=GG(c,d,a,b,x[k+7], S23,0x676F02D9);
		b=GG(b,c,d,a,x[k+12],S24,0x8D2A4C8A);
		a=HH(a,b,c,d,x[k+5], S31,0xFFFA3942);
		d=HH(d,a,b,c,x[k+8], S32,0x8771F681);
		c=HH(c,d,a,b,x[k+11],S33,0x6D9D6122);
		b=HH(b,c,d,a,x[k+14],S34,0xFDE5380C);
		a=HH(a,b,c,d,x[k+1], S31,0xA4BEEA44);
		d=HH(d,a,b,c,x[k+4], S32,0x4BDECFA9);
		c=HH(c,d,a,b,x[k+7], S33,0xF6BB4B60);
		b=HH(b,c,d,a,x[k+10],S34,0xBEBFBC70);
		a=HH(a,b,c,d,x[k+13],S31,0x289B7EC6);
		d=HH(d,a,b,c,x[k+0], S32,0xEAA127FA);
		c=HH(c,d,a,b,x[k+3], S33,0xD4EF3085);
		b=HH(b,c,d,a,x[k+6], S34,0x4881D05);
		a=HH(a,b,c,d,x[k+9], S31,0xD9D4D039);
		d=HH(d,a,b,c,x[k+12],S32,0xE6DB99E5);
		c=HH(c,d,a,b,x[k+15],S33,0x1FA27CF8);
		b=HH(b,c,d,a,x[k+2], S34,0xC4AC5665);
		a=II(a,b,c,d,x[k+0], S41,0xF4292244);
		d=II(d,a,b,c,x[k+7], S42,0x432AFF97);
		c=II(c,d,a,b,x[k+14],S43,0xAB9423A7);
		b=II(b,c,d,a,x[k+5], S44,0xFC93A039);
		a=II(a,b,c,d,x[k+12],S41,0x655B59C3);
		d=II(d,a,b,c,x[k+3], S42,0x8F0CCC92);
		c=II(c,d,a,b,x[k+10],S43,0xFFEFF47D);
		b=II(b,c,d,a,x[k+1], S44,0x85845DD1);
		a=II(a,b,c,d,x[k+8], S41,0x6FA87E4F);
		d=II(d,a,b,c,x[k+15],S42,0xFE2CE6E0);
		c=II(c,d,a,b,x[k+6], S43,0xA3014314);
		b=II(b,c,d,a,x[k+13],S44,0x4E0811A1);
		a=II(a,b,c,d,x[k+4], S41,0xF7537E82);
		d=II(d,a,b,c,x[k+11],S42,0xBD3AF235);
		c=II(c,d,a,b,x[k+2], S43,0x2AD7D2BB);
		b=II(b,c,d,a,x[k+9], S44,0xEB86D391);
		a=AddUnsigned(a,AA);
		b=AddUnsigned(b,BB);
		c=AddUnsigned(c,CC);
		d=AddUnsigned(d,DD);
	}
 
	var temp = WordToHex(a)+WordToHex(b)+WordToHex(c)+WordToHex(d);
 
	return temp.toLowerCase();
}`

	hexMD5 = `hexMD5 = function (e) {
  function i(e, i) {
    return e << i | e >>> 32 - i
  }
  function t(e, i) {
    var t,
    o,
    s,
    n,
    a;
    return s = 2147483648 & e,
    n = 2147483648 & i,
    t = 1073741824 & e,
    o = 1073741824 & i,
    a = (1073741823 & e) + (1073741823 & i),
    t & o ? 2147483648 ^ a ^ s ^ n : t | o ? 1073741824 & a ? 3221225472 ^ a ^ s ^ n : 1073741824 ^ a ^ s ^ n : a ^ s ^ n
  }
  function o(e, i, t) {
    return e & i | ~e & t
  }
  function s(e, i, t) {
    return e & t | i & ~t
  }
  function n(e, i, t) {
    return e ^ i ^ t
  }
  function a(e, i, t) {
    return i ^ (e | ~t)
  }
  function r(e, s, n, a, r, l, d) {
    return e = t(e, t(t(o(s, n, a), r), d)),
    t(i(e, l), s)
  }
  function l(e, o, n, a, r, l, d) {
    return e = t(e, t(t(s(o, n, a), r), d)),
    t(i(e, l), o)
  }
  function d(e, o, s, a, r, l, d) {
    return e = t(e, t(t(n(o, s, a), r), d)),
    t(i(e, l), o)
  }
  function c(e, o, s, n, r, l, d) {
    return e = t(e, t(t(a(o, s, n), r), d)),
    t(i(e, l), o)
  }
  function u(e) {
    for (var i, t = e.length, o = t + 8, s = (o - o % 64) / 64, n = 16 * (s + 1), a = Array(n - 1), r = 0, l = 0; t > l; ) i = (l - l % 4) / 4,
    r = l % 4 * 8,
    a[i] = a[i] | e.charCodeAt(l) << r,
    l++;
    return i = (l - l % 4) / 4,
    r = l % 4 * 8,
    a[i] = a[i] | 128 << r,
    a[n - 2] = t << 3,
    a[n - 1] = t >>> 29,
    a
  }
  function h(e) {
    var i,
    t,
    o = '',
    s = '';
    for (t = 0; 3 >= t; t++) i = e >>> 8 * t & 255,
    s = '0' + i.toString(16),
    o += s.substr(s.length - 2, 2);
    return o
  }
  function p(e) {
    e = e.replace(/\r\n/g, '\n');
    for (var i = '', t = 0; t < e.length; t++) {
      var o = e.charCodeAt(t);
      128 > o ? i += String.fromCharCode(o)  : o > 127 && 2048 > o ? (i += String.fromCharCode(o >> 6 | 192), i += String.fromCharCode(63 & o | 128))  : (i += String.fromCharCode(o >> 12 | 224), i += String.fromCharCode(o >> 6 & 63 | 128), i += String.fromCharCode(63 & o | 128))
    }
    return i
  }
  var g,
  m,
  f,
  v,
  w,
  C,
  b,
  k,
  y,
  F = Array(),
  T = 7,
  x = 12,
  P = 17,
  A = 22,
  S = 5,
  $ = 9,
  W = 14,
  B = 20,
  _ = 4,
  R = 11,
  I = 16,
  L = 23,
  G = 6,
  M = 10,
  O = 15,
  D = 21;
  for (e = p(e), F = u(e), C = 1732584193, b = 4023233417, k = 2562383102, y = 271733878, g = 0; g < F.length; g += 16) m = C,
  f = b,
  v = k,
  w = y,
  C = r(C, b, k, y, F[g + 0], T, 3614090360),
  y = r(y, C, b, k, F[g + 1], x, 3905402710),
  k = r(k, y, C, b, F[g + 2], P, 606105819),
  b = r(b, k, y, C, F[g + 3], A, 3250441966),
  C = r(C, b, k, y, F[g + 4], T, 4118548399),
  y = r(y, C, b, k, F[g + 5], x, 1200080426),
  k = r(k, y, C, b, F[g + 6], P, 2821735955),
  b = r(b, k, y, C, F[g + 7], A, 4249261313),
  C = r(C, b, k, y, F[g + 8], T, 1770035416),
  y = r(y, C, b, k, F[g + 9], x, 2336552879),
  k = r(k, y, C, b, F[g + 10], P, 4294925233),
  b = r(b, k, y, C, F[g + 11], A, 2304563134),
  C = r(C, b, k, y, F[g + 12], T, 1804603682),
  y = r(y, C, b, k, F[g + 13], x, 4254626195),
  k = r(k, y, C, b, F[g + 14], P, 2792965006),
  b = r(b, k, y, C, F[g + 15], A, 1236535329),
  C = l(C, b, k, y, F[g + 1], S, 4129170786),
  y = l(y, C, b, k, F[g + 6], $, 3225465664),
  k = l(k, y, C, b, F[g + 11], W, 643717713),
  b = l(b, k, y, C, F[g + 0], B, 3921069994),
  C = l(C, b, k, y, F[g + 5], S, 3593408605),
  y = l(y, C, b, k, F[g + 10], $, 38016083),
  k = l(k, y, C, b, F[g + 15], W, 3634488961),
  b = l(b, k, y, C, F[g + 4], B, 3889429448),
  C = l(C, b, k, y, F[g + 9], S, 568446438),
  y = l(y, C, b, k, F[g + 14], $, 3275163606),
  k = l(k, y, C, b, F[g + 3], W, 4107603335),
  b = l(b, k, y, C, F[g + 8], B, 1163531501),
  C = l(C, b, k, y, F[g + 13], S, 2850285829),
  y = l(y, C, b, k, F[g + 2], $, 4243563512),
  k = l(k, y, C, b, F[g + 7], W, 1735328473),
  b = l(b, k, y, C, F[g + 12], B, 2368359562),
  C = d(C, b, k, y, F[g + 5], _, 4294588738),
  y = d(y, C, b, k, F[g + 8], R, 2272392833),
  k = d(k, y, C, b, F[g + 11], I, 1839030562),
  b = d(b, k, y, C, F[g + 14], L, 4259657740),
  C = d(C, b, k, y, F[g + 1], _, 2763975236),
  y = d(y, C, b, k, F[g + 4], R, 1272893353),
  k = d(k, y, C, b, F[g + 7], I, 4139469664),
  b = d(b, k, y, C, F[g + 10], L, 3200236656),
  C = d(C, b, k, y, F[g + 13], _, 681279174),
  y = d(y, C, b, k, F[g + 0], R, 3936430074),
  k = d(k, y, C, b, F[g + 3], I, 3572445317),
  b = d(b, k, y, C, F[g + 6], L, 76029189),
  C = d(C, b, k, y, F[g + 9], _, 3654602809),
  y = d(y, C, b, k, F[g + 12], R, 3873151461),
  k = d(k, y, C, b, F[g + 15], I, 530742520),
  b = d(b, k, y, C, F[g + 2], L, 3299628645),
  C = c(C, b, k, y, F[g + 0], G, 4096336452),
  y = c(y, C, b, k, F[g + 7], M, 1126891415),
  k = c(k, y, C, b, F[g + 14], O, 2878612391),
  b = c(b, k, y, C, F[g + 5], D, 4237533241),
  C = c(C, b, k, y, F[g + 12], G, 1700485571),
  y = c(y, C, b, k, F[g + 3], M, 2399980690),
  k = c(k, y, C, b, F[g + 10], O, 4293915773),
  b = c(b, k, y, C, F[g + 1], D, 2240044497),
  C = c(C, b, k, y, F[g + 8], G, 1873313359),
  y = c(y, C, b, k, F[g + 15], M, 4264355552),
  k = c(k, y, C, b, F[g + 6], O, 2734768916),
  b = c(b, k, y, C, F[g + 13], D, 1309151649),
  C = c(C, b, k, y, F[g + 4], G, 4149444226),
  y = c(y, C, b, k, F[g + 11], M, 3174756917),
  k = c(k, y, C, b, F[g + 2], O, 718787259),
  b = c(b, k, y, C, F[g + 9], D, 3951481745),
  C = t(C, m),
  b = t(b, f),
  k = t(k, v),
  y = t(y, w);
  var z = h(C) + h(b) + h(k) + h(y);
  return z.toLowerCase()
},
recursiveReplaceDoublePointsBeforeLetter = function (e) {
  if ('object' == typeof e) for (var i in e) e[i] = recursiveReplaceDoublePointsBeforeLetter(e[i]);
   else if ('string' == typeof e) for (var t = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', i = 0; i < t.length; i++) {
    var o = new RegExp('\\.\\.' + t.charAt(i), 'g');
    e = e.replace(o, '.' + t.charAt(i))
  }
  return e
};`

	calcSig = `function CalcSig(e) {
  var i = [];
  for (key in e)
  	('string' == typeof e[key] || 'number' == typeof e[key] && e[key] % 1 === 0) && i.push(key);

  i.sort();

  for (var t = '', o = 0; o < i.length; o++)
  	t += (e[i[o]] + '').trim();

  return hexMD5(network + t + (sessionId + 1 << 5))
}`
)

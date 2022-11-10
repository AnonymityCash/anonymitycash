Anonymitycash version 1.1.0 is now available from:

  https://github.com/Anonymitycash/anonymitycash/releases/tag/v1.1.0


Please report bugs using the issue tracker at github:

  https://github.com/Anonymitycash/anonymitycash/issues

How to Upgrade
===============

Please notice new version anonymitycash path is $GOPATH/src/github.com/anonymitycash/anonymitycash if you install anonymitycash from source code.  
If you are running an older version, shut it down. Wait until it has quited completely, and then run the new version Anonymitycash.
You can operate according to the user manual.[(Anonymitycash User Manual)](https://anonymityswap.com/wp-content/themes/freddo/images/wallet/AnonymitycashUsermanualV1.0_en.pdf)


1.1.0 changelog
================
__Anonymitycash Node__

+ [`PR #1805`](https://github.com/Anonymitycash/anonymitycash/pull/1805)
    - Correct anonymitycash go import path to github.com/anonymitycash/anonymitycash. Developer can use go module to manage dependency of anonymitycash. 
+ [`PR #1815`](https://github.com/Anonymitycash/anonymitycash/pull/1815) 
    - Add asynchronous validate transactions function to optimize the performance of validating and saving block. 

__Anonymitycash Dashboard__

+ [`PR #1829`](https://github.com/Anonymitycash/anonymitycash/pull/1829) 
    - Fixed the decimals type string to integer in create asset page.

Credits
--------

Thanks to everyone who directly contributed to this release:

- DeKaiju
- iczc
- Paladz
- zcc0721
- ZhitingLin

And everyone who helped test.

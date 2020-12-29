# regex-rename

A simple program to rename files using regex. The regex dialect will match those of Go.

# sample

```
josh@akimi ~/syncthing/lectures/abstract-algebra $ regex-rename
regex find and regex replace are missing
  -c    Replace files that won't produce a duplicate, even when others will
  -continue-on-dupe
        Replace files that won't produce a duplicate, even when others will
  -dry-run
        Only print the renames
  -p    Only print the renames
  -s    Do not print the files repalced
  -silent
        Do not print the files repalced
josh@akimi ~/syncthing/lectures/abstract-algebra $ ls
'Lec 1 _ Abstract Algebra-aPKRuTDMguk.mp4'   'Lec 19 _ Abstract Algebra-mhDj_b5EXvY.mp4'  'Lec 29 _ Abstract Algebra-EN8QlPrmnAM.mp4'  'Lec 4 _ Abstract Algebra-9DId0iMTN6o.mp4'
'Lec 2 _ Abstract Algebra-vzzaZiaIMuk.mp4'  'Lec 20 _ Abstract Algebra-MWjl4KUdtWw.mp4'  'Lec 2 _ Abstract Algebra-e-a8auViFM0.mp4'   'Lec 5 _ Abstract Algebra-IlFpiGhJ9cA.mp4'
'Lec 11 _ Abstract Algebra-aPKRuTDMguk.mp4'  'Lec 21 _ Abstract Algebra-uG7zylsKp0s.mp4'  'Lec 30 _ Abstract Algebra-itvvxpntpwg.mp4'  'Lec 6 _ Abstract Algebra-B1rRzljv9Ac.mp4'
'Lec 12 _ Abstract Algebra-vzzaZiaIMuk.mp4'  'Lec 22 _ Abstract Algebra-mR9EbLL7uJk.mp4'  'Lec 31 _ Abstract Algebra-l1OWAasmBLI.mp4'  'Lec 7 _ Abstract Algebra-qYUBUU_LkFc.mp4'
'Lec 13 _ Abstract Algebra-pYAOHhS0WNA.mp4'  'Lec 23 _ Abstract Algebra-6NghC6eFAqg.mp4'  'Lec 32 _ Abstract Algebra-Ln6IgLksDxE.mp4'  'Lec 8 _ Abstract Algebra-bkZQ2drFHN4.mp4'
'Lec 14 _ Abstract Algebra-MgH9V0YBKpQ.mp4'  'Lec 24 _ Abstract Algebra-mVmN8ZI-ANg.mp4'  'Lec 33 _ Abstract Algebra-TsLWvWf4RdY.mp4'  'Lec 9 _ Abstract Algebra-SHjn798aWz4.mp4'
'Lec 15 _ Abstract Algebra-5iNwFTZYH2k.mp4'  'Lec 25 _ Abstract Algebra-YRdSQtGo73Y.mp4'  'Lec 34 _ Abstract Algebra-ACliYRhlvP8.mp4'
'Lec 16 _ Abstract Algebra-CZrnCopp6_M.mp4'  'Lec 26 _ Abstract Algebra-TXrOexBuODc.mp4'  'Lec 35 _ Abstract Algebra-IwxkWKgi8ao.mp4'
'Lec 17 _ Abstract Algebra-tx2yCz8MEvU.mp4'  'Lec 27 _ Abstract Algebra-ArVbqzjQtgU.mp4'  'Lec 36 _ Abstract Algebra-FCo8v2rzG9A.mp4'
'Lec 18 _ Abstract Algebra-IGvn8-P3AU0.mp4'  'Lec 28 _ Abstract Algebra-69x_Np21UB8.mp4'  'Lec 3 _ Abstract Algebra-mwcNETa0KFI.mp4'
josh@akimi ~/syncthing/lectures/abstract-algebra $ regex-rename "Lec (\d).+" "lec\0$1.mp4"
'Lec 1 _ Abstract Algebra-aPKRuTDMguk.mp4' -> 'lec01.mp4'
'Lec 2 _ Abstract Algebra-vzzaZiaIMuk.mp4' -> 'lec02.mp4'
'Lec 3 _ Abstract Algebra-mwcNETa0KFI.mp4' -> 'lec03.mp4'
'Lec 4 _ Abstract Algebra-MgH9V0YBKpQ.mp4' -> 'lec04.mp4'
'Lec 5 _ Abstract Algebra-5iNwFTZYH2k.mp4' -> 'lec05.mp4'
'Lec 6 _ Abstract Algebra-CZrnCopp6_M.mp4' -> 'lec06.mp4'
'Lec 7 _ Abstract Algebra-tx2yCz8MEvU.mp4' -> 'lec07.mp4'
'Lec 8 _ Abstract Algebra-IGvn8-P3AU0.mp4' -> 'lec08.mp4'
'Lec 9 _ Abstract Algebra-mhDj_b5EXvY.mp4' -> 'lec09.mp4'
josh@akimi ~/syncthing/lectures/abstract-algebra $ regex-rename -s "Lec (\d\d).+" "lec\$1.mp4"
josh@akimi ~/syncthing/lectures/abstract-algebra $ ls
lec01.mp4  lec04.mp4  lec07.mp4  lec10.mp4  lec13.mp4  lec16.mp4  lec19.mp4  lec22.mp4  lec25.mp4  lec28.mp4  lec31.mp4  lec34.mp4
lec02.mp4  lec05.mp4  lec08.mp4  lec11.mp4  lec14.mp4  lec17.mp4  lec20.mp4  lec23.mp4  lec26.mp4  lec29.mp4  lec32.mp4  lec35.mp4
lec03.mp4  lec06.mp4  lec09.mp4  lec12.mp4  lec15.mp4  lec18.mp4  lec21.mp4  lec24.mp4  lec27.mp4  lec30.mp4  lec33.mp4  lec36.mp4

- https://takeuforward.org/data-structure/longest-common-subsequence-dp-25/
- https://florian.github.io/diffing/
-https://medium.com/@snassr/processing-large-files-in-go-golang-6ea87effbfe2

- goroutines: https://dev.to/neelp03/how-to-use-goroutines-for-concurrent-processing-in-go-34ph

Todo:
- update generateDiff function to display deletedline addedline under each other
- update lcs to only use prev table calculation
- think about how to structure the diffing, should you return the diffing list or how you should do it
- don't forget to do the unit tests
- check how your function naming should be in Go
- how to create debug.json in VS Code
[[diffLine, Idx]]
[]
- lcs, unchanged, removed, inserted = lcs(text1[], text2[])
- diff, lineChangesTracker = generateDiff(text1[], text2[], removed, inserted)
- displayDiffWithContextLines(contextLinesDepth, diff, LineChangesTraacker, removed, inserted, text1, text2)
- I = 0
loop diffs:
    startIndex = diff[i][1]
    EndIndex   = diff[i][1]
         // calculate diff Width
        - for i = i+1 ; len(diff)-1:
            if IndexExist(diff[i+1][1], removed):
                endIndex   =  endIndex + 1
            
            if IndexExist(diff[i+1][1], inserted):
                endIndex   =  endIndex + 1

    
    I = EndIndex + 1

    for i:= startIndex - depth; i < endIndex + depth; i++ {
        if IndexExist(i, removed) || IndexExist(i, inserted):
           fmt.println(diff[i][0])
    }
    
    if 
                lineContext =  Diff[]   
    <!-- startIndexOftheCurrentChange
    endIndexofOftheCurrentChange
   - get indexof current diffLine
   - calculateCurrentDiffContextLines(curentDiffLineIndex, lineChangesTracker, depth)
        startIndex =   currentDiffLineIndex - depth
        endIndex   =  currentDiffLineIndex + depth
        
        if startIndex == -1:
           startIndex = 0

        if endIndex > len(text1) - 1:
           endIndex = len(text1) -1   
        i = 0
        for {
           if currentDiffLineIndex + 1 = 
        } -->

- loop through Text1
    


```diff text
1 unchanged line 1
2 unchanged line 2
3 unchanged line 3
4 - This is an example of a removed line.
5 + This is an example of an added line.
6 - Another line that was removed.
7 + Another line that was added.
8 unchanged line 4
9 unchanged line 5
10 unchanged line 6
11 - Another line that was removed.
12 + Another line that was added.
13 unchanged line 7
14 unchanged line 8
15 unchanged line 9
```

```diff
0 - line 1
1 line 2 
2 line 3
```


printDiff withContext solution.

case 1:
removed = [1,2,3,10,11]
inserted = [1,2,3,10,11]
changeTracker = [1,2,3,10,11]

case 2:
removed = [1,2]
inserted = []

case 3:
removed = []
inserted = [10,12,13]
changeTracker = [10,12,13]

case 4:
removed = [6,7,8]
inserted = [10,12,13]
changeTracker = [6,7,8,10,12,13]


diff = [- line,- line, - line ]



var changeStartIndex int
var changeEndIndex int
lastchangeIteratedIndex = -1

loop changeEndIndex > len(diff):

changeStartIndex = lastchangeIteratedIndex + 1
changeEndIndex  = lastchangeIteratedIndex + 1
// Calculate consecutive changes width
loop through changeTracker:
   if changeTracker[i] exist in removed and inserted:
   // modification
   changeEndIndex + 2

   if changeTracker[i] exist in removed and !exist in insert:
   // removed
   changeEndIndex + 1

   if changeTracker[i] !exist in removed and exist in insert::
   // insertions
   changeEndIndex + 1

   if changeTracker[i] != changeTracker[i+1]:
      lastchangeIteratedIndex = i
      break
      // finish the the consecutive change

// calculate start idx and end idx for context lines

ctxLineStartIdx = changeStartIndex - depth
ctxLineEndIdx = changeendIndex + depth

// Examine conextLinesIdx correctness

if ctxLineStartIdx < 0  : ctxLineStartIdx = 0
if ctxLineEndIdx > len(text2) : ctxLineEndIdx = len(text2)

// print the context changes
for i:= ctxLineStartIdx ; i < ctxLineEndIdx; i ++
   if changeStartIndex =< i && i <= changeEndIndex:
      fmt.frintln(diff[i])
   else:
      fmt.Println(text1[i])


case 1:
text1len = 1
text2len = 0
removed = [1]
inserted = []
changeTracker = [1]
text2Len = 5


file1 : [0,1,2,3]
file2 : [0,1,2,3]
case 1:
removed = [1,2,3,10,11]
inserted = [1,2,3,10,11]
changeTracker = [1,2,3,10,11]
text2Len = 5



case 2:
removed = [1,2]
inserted = []
changeTracker = [1,2]
text2Len = text1len - 2

case 3:
removed = []
inserted = [10,12,13]
changeTracker = [10,12,13]
text2Len = text1len + 3

case 4:
removed = [6,7,8]
inserted = [10,12,13]
changeTracker = [6,7,8,10,12,13]






diff = [{""- line 1"},{"+ line er"},"- line 2", "+ line fwef"]
removed = [0, 1]
inserted = [0, 1]
changesTracker = [0, 1]
ctxLineStartIdx = 0 
ctxLineEndIdx = 3

i = 2


for i in changesTracker:


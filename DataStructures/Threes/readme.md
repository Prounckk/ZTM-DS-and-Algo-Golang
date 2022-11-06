

## Balanced vs. Unbalanced
Balanced trees are trees where the height of every node's left and right subtrees differ by at most one. 
```
     6
   /   \
  4     8
 / \   / \ 
3   5 7   9
```

Unbalanced three is a linked list.
It becomes O(n) and slow to search if not balanced:
Example:
```
    8
   / \
  4   6
 / \   \ 
3   5   7
         \
          9
           \
            19
              \
               27
etc.       
 ```


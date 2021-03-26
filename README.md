# mutants-analyser-API

## Technical considerations

### Storage

We consider two different approaches for storing the data on a SQL database:

1. Storing only one table with only two columns one for the type *human* or *mutant* and one column for the count.
    The problem of this approach is that if the API is going to add new functionality, like getting all submitted dna,
    it won't be possible because we are only storing the count of each analysis result. But it is the most 
    memory and time effective approach. 

2. Storing each dna submitted and then when the stats are calculated, sum the amount of each type.
   This is the chosen approach because it says on the task that we need to save one register per DNA and it also will be able to scale,
    but it requires more memory and more time for each request.
   

### Analyser

For the analyser we implemented two types of analyser.
One allowing chains overlapping on the same direction and one that ignores them.
If we allow overlapping on same direction, any dna sequence containing a chain with at least five same nucleotides aligned on the same direction will be consider a mutant.
We can build two chains the one starting on the first nucleotide and ending in the fourth, and one starting on the second and ending on the fifth.
Also when we ignore chains overlapping on the same direction, we allow overlapping on different directions.
When you create the analyser you need to provide a boolean indicating if the analyser will consider or ignore chains overlapping on the same direction.


### Stats

When calculating mutant ratio, in the case of containing only mutants registers the ratio would be the quantity of mutants registered instead of infinity or error.
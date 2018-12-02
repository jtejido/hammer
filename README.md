# hammer

Heuristic Algorithms and Minimization Methods: An Expansive Repository

In artificial intelligence, an evolutionary algorithm (EA) is a subset of evolutionary computation.

In spite of the number (and it's rate of growth) and effectiveness of so and so "novel" population-based algorithms, the steps at which they usually work around at can be summarized by the following steps:

**Step One:** Generate the initial population of individuals randomly. (First generation)

**Step Two:** Evaluate the fitness of each individual in that population (iteration limit, sufficient fitness achieved, etc.)

**Step Three:** Repeat the following regenerational steps until termination:

1. Select the best-fit individuals for reproduction. (Parents)
2. Breed new individuals through crossover and mutation operations to give birth to offspring.
3. Evaluate the individual fitness of new individuals.
4. Replace least-fit population with new individuals.



This is an attempt at making a framework out of an otherwise duplicated approach of different EA algorithms (Others (i.e. swarm-based, quantum and simulated annealing, etc.) are still in my TO-DO list, and even though this has became a [problem](https://en.wikipedia.org/wiki/List_of_metaphor-based_metaheuristics#Criticism_of_the_metaphor_methodology) recently, their effectiveness is somewhat interesting to see in Information Retrieval). 

Hence, our acronym is HAMMER:

**-- if all you have is a hammer, everything looks like a nail.**


The approach is a set of interfaces that pretty much summarizes the above steps.

And therefore it is easy to separately inject each algorithms with a goal of minimizing and converging to a certain solution.

I will only add as many EAs and samples as deemed possible, as this was only a sub-component of an IR research project that I do (and I use this primarily as a Relevance feedback research tool using a float64 vector/map genetic sequence and Cosine Similarity/Jaccard Similarity as fitness function) and the framework helps me at generalizing (and ultimately helping me avoid re-typing) GA/GP/EP/DE's similar parts.


Install via:


```
go get github.com/jtejido/hammer/
```

Run the example at <GO_ROOT>/src/github.com/jtejido/hammer/examples/bitstring_ga_compare/ via :


```
go build
./bitstring_ga_compare
```

The Result would converge to an optimal solution like so (though in a real-world scenario, such a target is unknown, and EAs are used as a combinatorial optimization method to achieve an otherwise unknown goal, if it is even achievable given the time or iterations):




```
1 	 Bªhé¶z!eADåhCÙ=³á´û£Uôíö¨î9*%ì 	 175
2 	 òªLï:¤âqdnäô$SmÛa²ä¢UðíökzyjEÍÑ¯ 	 183
3 	 ¢(dü:&ðq
hWäç$e±Ka°ä£êÈ¥él¬äynEÍñ 	 189
4 	 â¨lm:¤pqdôä,KSm±èá°z!EÈ¤ml¨ä9äÅÕÑ 	 196
5 	 â¨dm:&ôqdêää-éSe±éa°ä!MÀ¥il)ä9oÅÄñ¯ 	 205
6 	 âhLm:¤ôqdêdå,éSu±éa°ä!EÀ¥ml¨ä9oGÅñ¯ 	 210
7 	 âhlm:¦ôqRdêdä,KCu±éa äùEÀ¥ml(ä9oÅÅñ¯ 	 213
8 	 âélm:&ôqZdêdå,KSu±Ya0ä!MÀ¤il¨ä9oGÅq 	 217
9 	 âilm:$ôqRdêäå,KSu±éa äùEÀ¤ml(ä9oGÅq¯ 	 221
10 	 âilm>&ôqdêdå,KSu±éa0äùMÀ¤il(ä9oGÅq¯ 	 223
11 	 âilm>$ôqRdêdå,KCe±éa kùEÁ¤ml(z9oGÅq¯ 	 226
12 	 âilm?&ôqRdêdå,KSe±éa kùMÀ¤ml(z9oGÅq/ 	 229
13 	 âilm?&ôqRdêdå,KSe1éa äùMÁ¤al(z9oGÅs/ 	 231
14 	 âilm>$ôqRdêdå,KSe1éa k+MÁ¤il(z9oGÅs/ 	 234
15 	 âilm?$ôqRdêdå,KSe1éa k+MÁ il(z9oGÅs/ 	 236
16 	 âilm?$ôqRdêlå,KSe1éa c+MÁ il(z9oGÅs/ 	 238
17 	 bilm?$ôqRdêlå,KSe1éa c+MÁ il(x9oGÅw/ 	 239
18 	 bilm?$ôqRdêlå,CSe1éa ckMÁ il(x9oGÅs/ 	 242
19 	 bilm?$ôqRdêlå,CSa1éa ckMÁ il(x9oGÅs/ 	 243
20 	 bilm?$ôqRdêlå,CSa1éa ckMÁ il(x9gGÅs/ 	 244
21 	 ôilm? ôqRdêlå,CSa1éa ckMÁ il(x9gGÅs/ 	 246
22 	 ôilm? ôqrdülå,CSa1éa ckMÁ il(x9gGÅs/ 	 249
23 	 ôilm? ôqrdölå,CSa1éa ckMÁ il(x9gGÅs/ 	 251
24 	 ôilm? àqrdölå,CSa1éa ckMÁ il(x9gGÅs/ 	 252
25 	 ôilm? àqrdölå.CSa1éa ckMÁ il(xygGÅs/ 	 254
26 	 ôilm? àqrdxlå.CSa1éa ckMÁ il(xygGEs/ 	 256
27 	 ôilm/ àqrdxlå.CSa1éa ckMÁ il(xyeGEs/ 	 258
28 	 ôilm/ àqrdxlå.CSa1Sa ckmÁ il(xyeGEs/ 	 260
29 	 ôilm/ àqrdxlå.CS!1Sa ckmÅ il(xyeGEs/ 	 262
30 	 ôall/ àqrdxlå.CS!1Sa ckmÅ il(xyeGEs/ 	 264
31 	 ôall/ àqrdxlå.cS!1Sa ckmÅ il(xieGEs/ 	 265
32 	 àall/ àqrdxlå.gS!1Sa ckmÅ il(xieGEs/ 	 267
33 	 àall/ àqrdxlí.Gs!1Sa ckmÅ il(xieGEs/ 	 269
34 	 àall/ àqrdxlí.gs!1Sa ckmÅ il(xieCEs/ 	 271
35 	 àall/ àqrdxlíngs!1Sa ckmÅ il(xieCEs/ 	 272
36 	 àall/ Dqrdxlí.gs!!Sa ckmÅ il(xieCEs/ 	 273
37 	 àall/ Dqrdxlíngs!!Sa ckmÅ il xieCEs/ 	 275
38 	 àallo Dqrtxlíngs!!Sa ckmÅ il(xieCEs/ 	 276
39 	 àallo Dqrtxlíngs!!Sa comÅ il xieCEs/ 	 278
40 	 àello Dqrdxlíngs!!Sa comÅ in pieCEs/ 	 280
41 	 àallo Eqrtxlíngs!!Sa comÅ in piecEs. 	 283
42 	 àello Eqrtxlíngs! Sa comÅ in piecEs/ 	 284
43 	 àello EqrTxlíngs! Se comÅ in piecEs. 	 285
44 	 àullo Eqrtxlíngs! Se comå in piecEs. 	 286
45 	 àello Eqrthlífgs! Se comå in piecEs. 	 287
46 	 àello Eqrtxlmngs! Se comå in pieces. 	 289
47 	 Hello Eqrtxlmngs! Se comå in pieces. 	 290
48 	 àello Eqrthlmngs! We comå in pieces. 	 291
49 	 Hello Eqrtxlmngs! We comå in pieces. 	 291
50 	 àello Eqrthlings! We comå in pieces. 	 292
51 	 Hello Eqrthlings! We come in"pieces. 	 293
52 	 Hello Eqrthlkngs! We come in pieces.. 	 294
53 	 Hello Eqrthlings! We come in pieces.. 	 295
54 	 Hello Eqrthlings! We come in pieces.. 	 295
55 	 Hello Earthlings! We come in pieces. 	 295
56 	 Hello Earthlings! We come in pieces.. 	 296
1.154213734s
Alloc = 3 MiB	TotalAlloc = 337 MiB	Sys = 12 MiB	NumGC = 111
```

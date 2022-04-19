# Sci-Go

## The Go Data Science Library

<img src="./resources/Gopher.png" width=400px>

This library is a data science library, not necessarily a Machine Learning or Deep Learning library. SciGo has been developed with an idiomatic design process, and inspiration from a few sources.

* [Go Proverbs With Rob Pike](https://www.youtube.com/watch?v=PAAkCSZUG1c)

## Why Should I Use SciGo?

1. SciGo doesn't lock you in
   * Inputs are almost always native Golang types, allowing you to interface with any awesome golang library like [GoDataframe](REFERENCE) TODO ref
   * We rely on CGo as little as possible (TODO, ref on how CGo causes lots of problems), yet have good performance (TODO, back this up)
2. SciGo is very easy to use
   * SciGo has a very clear and organized design pattern
   * We painstakingly limit our dependencies to make your life easier.
3. SciGo is ambitious and has a solid roadmap for a ton of features

## Installing

```bash
go get -u github.com/deleomike/sci-go
```

## Usage

There are a lot of things in the [works](https://github.com/deleomike/sci-go/issues/1), but right now we have these feature sets implemented.

* [Clustering](./cluster/README.md)
    * [KMeans](./cluster/README.md#kmeans)
    * [Mean Shift](./cluster/README.md#mean-shift)

## Support the Project

I appreciate all contributions and sponsorships for this project.

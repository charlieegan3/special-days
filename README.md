# special-days

This is a package to give dates for [Mothering
Sunday](https://en.wikipedia.org/wiki/Mothering_Sunday) and [Fathers Day
(UK)](https://en.wikipedia.org/wiki/Father%27s_Day#United_Kingdom).

```
go get github.com/charlieegan3/special-days
```

```
fathersDay, err :=  fathersday.FathersDay("uk", 2021) // => time.Time, err
```

```
motheringSunday, err := motheringsunday.MotheringSunday(2021) // => time.Time, err
```

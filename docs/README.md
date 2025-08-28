# Insights

## Overview

```vega
{
  "$schema": "https://vega.github.io/schema/vega-lite/v6.json",
  "width": 800,
  "height": 300,
  "title": "Top 10 Repositories by Visitors",
  "data": {
    "values": [
      {"repository": "vk_ads_test", "views": 16}
    ]
  },
  "mark": "bar",
  "encoding": {
    "y": {"field": "repository", "type": "nominal", "title": "Repository", "sort": "-x"},
    "x": {"field": "views", "type": "quantitative", "title": "Total Views"}
  }
}
```


```vega
{
  "$schema": "https://vega.github.io/schema/vega-lite/v6.json",
  "width": 800,
  "height": 300,
  "title": "Top 10 Repositories by Git Clones",
  "data": {
    "values": [
      {"repository": "vk_ads_test", "clones": 108}
    ]
  },
  "mark": "bar",
  "encoding": {
    "y": {"field": "repository", "type": "nominal", "title": "Repository", "sort": "-x"},
    "x": {"field": "clones", "type": "quantitative", "title": "Total Clones"}
  }
}
```

## Repository Breakdown

### source-Alexander-Rudenko/vk_ads_test

```vega
{
  "$schema": "https://vega.github.io/schema/vega-lite/v6.json",
  "width": 800,
  "title": "Visitors for source-Alexander-Rudenko/vk_ads_test",
  "data": {
    "values": [
      {"date": "2025-08-03", "type": "Total Views", "value": 16},
      {"date": "2025-08-03", "type": "Unique Views", "value": 1}
    ]
  },
  "mark": "line",
  "encoding": {
    "x": {
      "field": "date",
      "type": "temporal",
      "title": "Date",
      "scale": { "type": "utc" },
      "axis": {
        "format": "%Y-%m-%d",
        "labelAngle": -45,
        "labelOverlap": false,
        "tickCount": {"interval": "day", "step": 1}
      }
    },
    "y": {"field": "value", "type": "quantitative", "title": "Views"},
    "color": {
      "field": "type",
      "type": "nominal",
      "legend": {
        "title": null
      }
    },
    "tooltip": [
      { "field": "date", "type": "temporal", "title": "Date" },
      { "field": "type", "type": "nominal", "title": "Metric" },
      { "field": "value", "type": "quantitative", "title": "Value" }
    ]
  }
}
```


```vega
{
  "$schema": "https://vega.github.io/schema/vega-lite/v6.json",
  "width": 800,
  "title": "Git Clones for source-Alexander-Rudenko/vk_ads_test",
  "data": {
    "values": [
      {"date": "2025-07-09", "type": "Total Clones", "value": 2},
      {"date": "2025-07-10", "type": "Total Clones", "value": 2},
      {"date": "2025-07-11", "type": "Total Clones", "value": 2},
      {"date": "2025-07-12", "type": "Total Clones", "value": 2},
      {"date": "2025-07-13", "type": "Total Clones", "value": 3},
      {"date": "2025-07-14", "type": "Total Clones", "value": 3},
      {"date": "2025-07-15", "type": "Total Clones", "value": 4},
      {"date": "2025-07-16", "type": "Total Clones", "value": 3},
      {"date": "2025-07-17", "type": "Total Clones", "value": 4},
      {"date": "2025-07-18", "type": "Total Clones", "value": 2},
      {"date": "2025-07-19", "type": "Total Clones", "value": 2},
      {"date": "2025-07-20", "type": "Total Clones", "value": 2},
      {"date": "2025-07-21", "type": "Total Clones", "value": 4},
      {"date": "2025-07-22", "type": "Total Clones", "value": 1},
      {"date": "2025-07-23", "type": "Total Clones", "value": 1},
      {"date": "2025-07-24", "type": "Total Clones", "value": 2},
      {"date": "2025-07-25", "type": "Total Clones", "value": 4},
      {"date": "2025-07-26", "type": "Total Clones", "value": 3},
      {"date": "2025-07-27", "type": "Total Clones", "value": 3},
      {"date": "2025-07-28", "type": "Total Clones", "value": 1},
      {"date": "2025-07-29", "type": "Total Clones", "value": 3},
      {"date": "2025-07-30", "type": "Total Clones", "value": 3},
      {"date": "2025-07-31", "type": "Total Clones", "value": 1},
      {"date": "2025-08-01", "type": "Total Clones", "value": 1},
      {"date": "2025-08-02", "type": "Total Clones", "value": 2},
      {"date": "2025-08-03", "type": "Total Clones", "value": 1},
      {"date": "2025-08-04", "type": "Total Clones", "value": 2},
      {"date": "2025-08-05", "type": "Total Clones", "value": 1},
      {"date": "2025-08-06", "type": "Total Clones", "value": 1},
      {"date": "2025-08-07", "type": "Total Clones", "value": 4},
      {"date": "2025-08-08", "type": "Total Clones", "value": 1},
      {"date": "2025-08-09", "type": "Total Clones", "value": 2},
      {"date": "2025-08-10", "type": "Total Clones", "value": 1},
      {"date": "2025-08-11", "type": "Total Clones", "value": 1},
      {"date": "2025-08-12", "type": "Total Clones", "value": 3},
      {"date": "2025-08-13", "type": "Total Clones", "value": 1},
      {"date": "2025-08-14", "type": "Total Clones", "value": 4},
      {"date": "2025-08-15", "type": "Total Clones", "value": 1},
      {"date": "2025-08-16", "type": "Total Clones", "value": 4},
      {"date": "2025-08-17", "type": "Total Clones", "value": 1},
      {"date": "2025-08-18", "type": "Total Clones", "value": 2},
      {"date": "2025-08-19", "type": "Total Clones", "value": 3},
      {"date": "2025-08-20", "type": "Total Clones", "value": 1},
      {"date": "2025-08-21", "type": "Total Clones", "value": 2},
      {"date": "2025-08-22", "type": "Total Clones", "value": 1},
      {"date": "2025-08-23", "type": "Total Clones", "value": 2},
      {"date": "2025-08-24", "type": "Total Clones", "value": 2},
      {"date": "2025-08-25", "type": "Total Clones", "value": 3},
      {"date": "2025-08-26", "type": "Total Clones", "value": 1},
      {"date": "2025-08-27", "type": "Total Clones", "value": 3},
      {"date": "2025-07-09", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-10", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-11", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-12", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-13", "type": "Unique Clones", "value": 2},
      {"date": "2025-07-14", "type": "Unique Clones", "value": 2},
      {"date": "2025-07-15", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-16", "type": "Unique Clones", "value": 2},
      {"date": "2025-07-17", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-18", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-19", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-20", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-21", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-22", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-23", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-24", "type": "Unique Clones", "value": 2},
      {"date": "2025-07-25", "type": "Unique Clones", "value": 4},
      {"date": "2025-07-26", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-27", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-28", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-29", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-30", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-31", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-01", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-02", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-03", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-04", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-05", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-06", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-07", "type": "Unique Clones", "value": 4},
      {"date": "2025-08-08", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-09", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-10", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-11", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-12", "type": "Unique Clones", "value": 3},
      {"date": "2025-08-13", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-14", "type": "Unique Clones", "value": 4},
      {"date": "2025-08-15", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-16", "type": "Unique Clones", "value": 4},
      {"date": "2025-08-17", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-18", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-19", "type": "Unique Clones", "value": 3},
      {"date": "2025-08-20", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-21", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-22", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-23", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-24", "type": "Unique Clones", "value": 2},
      {"date": "2025-08-25", "type": "Unique Clones", "value": 3},
      {"date": "2025-08-26", "type": "Unique Clones", "value": 1},
      {"date": "2025-08-27", "type": "Unique Clones", "value": 3}
    ]
  },
  "mark": "line",
  "encoding": {
    "x": {
      "field": "date",
      "type": "temporal",
      "title": "Date",
      "scale": { "type": "utc" },
      "axis": {
        "format": "%Y-%m-%d",
        "labelAngle": -45,
        "labelOverlap": false,
        "tickCount": {"interval": "day", "step": 1}
      }
    },
    "y": {"field": "value", "type": "quantitative", "title": "Clones"},
    "color": {
      "field": "type",
      "type": "nominal",
      "legend": {
        "title": null
      }
    },
    "tooltip": [
      { "field": "date", "type": "temporal", "title": "Date" },
      { "field": "type", "type": "nominal", "title": "Metric" },
      { "field": "value", "type": "quantitative", "title": "Value" }
    ]
  }
}
```

| Referral Source | Views | Unique Visitors |
|-|-|-|


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
      {"repository": "vk_ads_test", "views": 8}
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
      {"repository": "vk_ads_test", "clones": 84}
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
      {"date": "2025-06-13", "type": "Total Views", "value": 2},
      {"date": "2025-06-14", "type": "Total Views", "value": 0},
      {"date": "2025-06-15", "type": "Total Views", "value": 3},
      {"date": "2025-06-16", "type": "Total Views", "value": 0},
      {"date": "2025-06-17", "type": "Total Views", "value": 0},
      {"date": "2025-06-18", "type": "Total Views", "value": 0},
      {"date": "2025-06-19", "type": "Total Views", "value": 0},
      {"date": "2025-06-20", "type": "Total Views", "value": 0},
      {"date": "2025-06-21", "type": "Total Views", "value": 3},
      {"date": "2025-06-13", "type": "Unique Views", "value": 1},
      {"date": "2025-06-14", "type": "Unique Views", "value": 0},
      {"date": "2025-06-15", "type": "Unique Views", "value": 1},
      {"date": "2025-06-16", "type": "Unique Views", "value": 0},
      {"date": "2025-06-17", "type": "Unique Views", "value": 0},
      {"date": "2025-06-18", "type": "Unique Views", "value": 0},
      {"date": "2025-06-19", "type": "Unique Views", "value": 0},
      {"date": "2025-06-20", "type": "Unique Views", "value": 0},
      {"date": "2025-06-21", "type": "Unique Views", "value": 1}
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
      {"date": "2025-06-10", "type": "Total Clones", "value": 8},
      {"date": "2025-06-11", "type": "Total Clones", "value": 5},
      {"date": "2025-06-12", "type": "Total Clones", "value": 0},
      {"date": "2025-06-13", "type": "Total Clones", "value": 0},
      {"date": "2025-06-14", "type": "Total Clones", "value": 0},
      {"date": "2025-06-15", "type": "Total Clones", "value": 22},
      {"date": "2025-06-16", "type": "Total Clones", "value": 4},
      {"date": "2025-06-17", "type": "Total Clones", "value": 3},
      {"date": "2025-06-18", "type": "Total Clones", "value": 4},
      {"date": "2025-06-19", "type": "Total Clones", "value": 2},
      {"date": "2025-06-20", "type": "Total Clones", "value": 2},
      {"date": "2025-06-21", "type": "Total Clones", "value": 3},
      {"date": "2025-06-22", "type": "Total Clones", "value": 2},
      {"date": "2025-06-23", "type": "Total Clones", "value": 3},
      {"date": "2025-06-24", "type": "Total Clones", "value": 2},
      {"date": "2025-06-25", "type": "Total Clones", "value": 2},
      {"date": "2025-06-26", "type": "Total Clones", "value": 2},
      {"date": "2025-06-27", "type": "Total Clones", "value": 2},
      {"date": "2025-06-28", "type": "Total Clones", "value": 2},
      {"date": "2025-06-29", "type": "Total Clones", "value": 4},
      {"date": "2025-06-30", "type": "Total Clones", "value": 4},
      {"date": "2025-07-01", "type": "Total Clones", "value": 4},
      {"date": "2025-07-02", "type": "Total Clones", "value": 2},
      {"date": "2025-07-03", "type": "Total Clones", "value": 2},
      {"date": "2025-06-10", "type": "Unique Clones", "value": 8},
      {"date": "2025-06-11", "type": "Unique Clones", "value": 5},
      {"date": "2025-06-12", "type": "Unique Clones", "value": 0},
      {"date": "2025-06-13", "type": "Unique Clones", "value": 0},
      {"date": "2025-06-14", "type": "Unique Clones", "value": 0},
      {"date": "2025-06-15", "type": "Unique Clones", "value": 14},
      {"date": "2025-06-16", "type": "Unique Clones", "value": 3},
      {"date": "2025-06-17", "type": "Unique Clones", "value": 2},
      {"date": "2025-06-18", "type": "Unique Clones", "value": 3},
      {"date": "2025-06-19", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-20", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-21", "type": "Unique Clones", "value": 2},
      {"date": "2025-06-22", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-23", "type": "Unique Clones", "value": 2},
      {"date": "2025-06-24", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-25", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-26", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-27", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-28", "type": "Unique Clones", "value": 1},
      {"date": "2025-06-29", "type": "Unique Clones", "value": 3},
      {"date": "2025-06-30", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-01", "type": "Unique Clones", "value": 3},
      {"date": "2025-07-02", "type": "Unique Clones", "value": 1},
      {"date": "2025-07-03", "type": "Unique Clones", "value": 1}
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


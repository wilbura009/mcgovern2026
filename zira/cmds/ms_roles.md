```json
z/embed {
  "color": 0,
  "title": "Role Selection",
  "description": "Select your MS status. Click one of the following reaction emojis to self-assign a role.",
  "fields": [
    {
      "name": "MS1 :one:",
      "value": "First-Year Medical Student.",
      "inline": true
    },
    {
      "name": "MS2 :two:",
      "value": "Second-Year Medical Student.",
      "inline": true
    }
  ],
  "footer": {
    "text": "Kindly let us know if a role is missing or a description is incorrect and we will work to fix it."
  }
}
```


```txt
z/add :one: @MS1
z/add :two: @MS2
```

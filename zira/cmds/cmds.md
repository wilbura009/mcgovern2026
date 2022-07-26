## Embedded Message

```json
z/embed {
  "color": 5683084,
  "title": "Role Selection",
  "description": "This section allows you to self-assign roles. Click one of the following emojis to assign yourself a role.",
  "fields": [
    {
      "name": "md-mph :blue_square:",
      "value": "---",
      "inline": true
    },
    {
      "name": "md-mba :yellow_square:",
      "value": "---",
      "inline": true
    },
    {
      "name": "md-phd :purple_square:",
      "value": "---",
      "inline": true
    },
    {
      "name": "md-omfs :orange_square:",
      "value": "---",
      "inline": true
    },
    {
      "name": "md-ms-biomedical-informatics :green_square:",
      "value": "---",
      "inline": true
    },
    {
      "name": "md-mbe :white_medium_square:",
      "value": "---",
      "inline": true
    },
    {
      "name": "hpsp :red_square:",
      "value": "---",
      "inline": true
    }
  ],
  "footer": {
    "text": "Kindly let us know if a role is missing or a description is incorrect and we will work to fix it."
  }
}
```

## Add Reaction Buttons

```txt
z/add :blue_square: @md-mph
z/add :yellow_square: @md-mba
z/add :purple_square: @md-phd
z/add :orange_square: @md-omfs
z/add :green_square: @md-biomedical
z/add :white_medium_square: @md-mbe
z/add :red_square: @hpsp
```

---

## Messages

### Welcome

```json
z/embed {
  "color": 0,
  "title": "Welcome to the Server!",
  "description": "Hello everyone!\n\nWhen you join, please change your server name to your actual name so we can know who’s who. The fastest way to do this would be to tap on the people button in the top right corner (:busts_in_silhouette:), tap your name on the list, and then go to “Edit Server Profile” to enter your name. Be sure to click the “Save” button afterwards too\n\nSince we have a lot more faces joining here, welcome again everyone!! Our server is still in the process of being organized as our needs for it evolve! I created a \"suggestions\" channel under “Support” if y’all have ideas on how to arrange it or if y’all want to add another channel for us! Some current MS1’s have been giving me some ideas too. So far, I have categories for school (for when it starts) and our interests (since I’ve seen a few of us talk about working out ayyy), etc.\n\nFor those whom are using Discord for the first time here, feel free to reach out if you have any questions! It is pretty intuitive; for ex, you can choose to mute certain channels that might not interest or apply to you. ",
  "fields": [
    {
      "name": "Suggestions",
      "value": "https://discord.com/channels/959642789409337374/960315427588812881",
      "inline": true
    },
    {
      "name": "Role Selection",
      "value": "https://discord.com/channels/959642789409337374/1000882010769862716",
      "inline": true
    }
  ],
  "footer": {
    "text": "Looking forward to meeting you all soon!! Hopefully we can get everyone on the server so that it can help us help each other "
  }
}
```

### Rules

#### 1 - 4

```json
z/embed {
  "color": 0,
  "title": "Rules",
  "description": "Please enjoy the community as courteous and responsible adults. This server exists to help us aid each other throughout the next four years together. Please read the following guidelines:",
  "fields": [
    {
      "name": "1. Be a decent human being",
      "value": "🔸 Be Kind and Respectful.\n🔸 Use the Golden Rule: treat others how you want to be treated.\n🔸 Derogatory language towards any user will not be tolerated.",
      "inline": false
    },
    {
      "name": "2. Trigger Topics",
      "value": "🔸 Be courteous of other members and provide a warning for topics you may deem sensitive/controversial.",
      "inline": false
    },
    {
      "name": "3. Intolerance",
      "value": "🔸 Intolerance of any kind is not allowed.\n🔸 Do NOT use slurs or phrases that are offensive to any group of people.\n🔸 Absolutely NO racism, homophobia, transphobia, misogyny, xenophobia, anti-semitism, or otherwise discriminatory behavior/language will be tolerated.\n🔸 Absolutely NO NSFW content is allowed within the server. ",
      "inline": false
    },
    {
      "name": "4. Politics",
      "value": "🔸 Keep political conversations to a minimum. Everyone is entitled to their opinions, but we want to keep tension out of this server.\n🔸 Feel free to have political discourse through other platforms. ",
      "inline": false
    }
  ]
}
```

#### 5 - 7

```json
z/embed {
  "color": 0,
  "fields": [
    {
      "name": "5. Keep channels on topic",
      "value": "🔸 Find the appropriate channel for your discussion.\n🔸 We are aware that conversations flow organically so do not feel pressure to adhere strictly to this rule.\n",
      "inline": false
    },
    {
      "name": "6. Spam",
      "value": "🔸 No spamming messages in the server or to the community members via DMs.\n🔸 Do not @ everyone in your messages. If you believe it is absolutely necessary, message one of the moderators.\n",
      "inline": false
    },
    {
      "name": "7.  Have any questions?",
      "value": "🔸 The moderators are here to help maintain the server and adjust to the needs of the community. If you have any questions, comments, ideas, or concerns, please contact the moderators via DM or through the suggestions channel.\n🔸 Moderators: Courtney N, Rosa Burroughs, & Nikhil Vj",
      "inline": false
    },
    {
       "name": "Compliance",
       "value": "🔹 Your presence in this server implies accepting these rules, including all further changes. These changes might be done at any time without notice; please be sure to stay up to date on guidelines.",
       "inline": false
    }
  ],
  "footer": {
    "text": "We hope to maintain harmony in our community:v_tone3:. \nAll issues in the community will be addressed and resolved."
  }
}
```

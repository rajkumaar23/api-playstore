# Unofficial PlayStore API

**Explore android app metadata from PlayStore with ease!**

[![build](https://github.com/rajkumaar23/api-playstore/actions/workflows/build.yaml/badge.svg)](https://github.com/rajkumaar23/api-playstore/actions/workflows/build.yaml)

## Usage

Retrieve comprehensive JSON data for any app from Google Play by making a simple GET request:
```
https://api-playstore.rajkumaar.co.in/json?id=<YOUR-PACKAGE-ID>
```
Example : https://api-playstore.rajkumaar.co.in/json?id=com.openai.chatgpt

Sample JSON response with all possible details:
```json
{
    "packageID": "com.openai.chatgpt",
    "name": "ChatGPT",
    "version": "1.2023.319",
    "downloads": "10,000,000+",
    "downloadsExact": 40282110,
    "lastUpdated": "Dec 8, 2023",
    "launchDate": "Jul 21, 2023",
    "developer": "OpenAI",
    "description": "OpenAI’s latest advancements at your fingertips.\u003cbr\u003e\u003cbr\u003eThis official app is free, syncs your history across devices, and brings you the newest model improvements from OpenAI.\u003cbr\u003e\u003cbr\u003eWith ChatGPT in your pocket, you’ll find:\u003cbr\u003e\u003cbr\u003e· Instant answers\u003cbr\u003e· Tailored advice\u003cbr\u003e· Creative inspiration\u003cbr\u003e· Professional input\u003cbr\u003e· Learning opportunities\u003cbr\u003e\u003cbr\u003eJoin millions of users and try out the app that’s been captivating the world. Download ChatGPT today.\u003cbr\u003e\u003cbr\u003eTerms of service \u0026amp; privacy policy:\u003cbr\u003ehttps://openai.com/policies/terms-of-use\u003cbr\u003ehttps://openai.com/policies/privacy-policy",
    "screenshots": [
        "https://play-lh.googleusercontent.com/73YdqpsZPOjPdiDrURm2whbE-CAoIvPLPSpoH76y4vDz-K19JDIutwWBiHWY-1SzACs",
        "https://play-lh.googleusercontent.com/bw67GNVLWhLKv8taCwdyttVzmI0R-NVHOZ6ms33Cpz1UVNn9ZAR_B2E_PbtIeouKDoQ",
        "https://play-lh.googleusercontent.com/pfSjG93dd5Xy2DBh4FsHJPFNWz5sXB1Dwg_qPzeXXw_P7Mq6oJxQz8FrdW1dEj8K6IXN",
        "https://play-lh.googleusercontent.com/8Z4_BoVmJkSEftmXGpW3Fzdstq9DZKO55BrOfDhpOuPTV3Hdm5ADdTAQ3LIKmE9uwUj6",
        "https://play-lh.googleusercontent.com/eXCnYvgm8DAyVaBTKTQoE-zhhvI1VcsuCunsEOtmZmjkFtrNoC1GqoXb4zeFbCJQgw",
        "https://play-lh.googleusercontent.com/jYSdAP5tqMrku4A9P3bHbqG7vnuQOtpAfzMJ_y7sGDusu4Bd-myy0A5eOZYYj3D73sc"
    ],
    "category": "Productivity",
    "logo": "https://play-lh.googleusercontent.com/6qi3w4uqKaD1c-CBdkkfO6IL0lH4OoCTEdiX0oYbLFxwfvxu1t8vuwHcagdYSFmFKmI",
    "banner": "https://play-lh.googleusercontent.com/ZUHDpTlKqnmnqQJTgJIy2hdrYy0oqhF7v3pbjMcoYDjBr843HxPzQnvZU6TczCZPRwg",
    "privacy_policy": "https://openai.com/privacy",
    "latest_update_message": "",
    "website": "https://help.openai.com",
    "support_email": "support@openai.com",
    "rating": "4.8",
    "noOfUsersRated": "318,000"
}
```

## Integration with [Shields.io](https://shields.io)
Customize your README with dynamic badges for specific attributes. Utilize the following template:
```markdown
![Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/<ATTRIBUTE-NAME>?id=<PACKAGE-ID>)
```
Supported attributes: 
- packageID
- name
- version
- downloads
- downloadsExact
- lastUpdated
- launchDate
- developer
- description
- screenshots
- category
- logo
- banner
- privacy_policy
- latest_update_message
- website
- support_email
- rating
- noOfUsersRated

Explore customization options based on shields.io style parameters. Learn more [here](https://shields.io/).

### Examples

Downloads
```markdown
![Downloads Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/downloads?id=com.openai.chatgpt)
```
![Downloads Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/downloads?id=com.openai.chatgpt)

Version
```markdown
![Version Badge](https://img.shields.io/endpoint?color=blue&url=https://api-playstore.rajkumaar.co.in/version?id=com.openai.chatgpt)
```
![Version Badge](https://img.shields.io/endpoint?color=blue&url=https://api-playstore.rajkumaar.co.in/version?id=com.openai.chatgpt)

Developer
```markdown
![Dev Badge](https://img.shields.io/endpoint?color=orange&style=for-the-badge&url=https://api-playstore.rajkumaar.co.in/developer?id=com.openai.chatgpt)
```
![Dev Badge](https://img.shields.io/endpoint?color=orange&style=for-the-badge&url=https://api-playstore.rajkumaar.co.in/developer?id=com.openai.chatgpt)

### License
Find the license [here](LICENSE)



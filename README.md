# Unofficial PlayStore API

**Explore android app metadata from PlayStore with ease!**

[Available on RapidAPI](https://rapidapi.com/rajkumaar23-RK1kHf2Zl/api/app-details-from-playstore/details)

## Usage

Retrieve comprehensive JSON data for any app from Google Play by making a simple GET request:
```
https://api-playstore.rajkumaar.co.in/json?id=<YOUR-PACKAGE-ID>
```
Example : https://api-playstore.rajkumaar.co.in/json?id=in.co.rajkumaar.amritarepo

Sample JSON response with all possible details:
```json
{
  "packageID": "in.co.rajkumaar.amritarepo",
  "name": "Amrita Repository",
  "version": "4.3.13",
  "downloads": "10,000+",
  "downloadsExact": 14105,
  "lastUpdated": "Nov 12, 2023",
  "launchDate": "Jul 22, 2018",
  "developer": "Rajkumar S",
  "description": "THE-LONG-APP-DESCRIPTION",
  "screenshots": [
    "https://play-lh.googleusercontent.com/EEZ8RKfIYtedPSVLH24hjAeBkGx3vkZCksypXkmJypGb88ZbmnSNI20fHlti1q5q9dXa",
    "https://play-lh.googleusercontent.com/4y1_wfSrMxiKd54i8pc8AuvmCdD_Wqh7zu4kRcJoW4VWfoXOflb-KZzHPTjhWMmkptsb",
  ],
  "category": "Education",
  "logo": "https://play-lh.googleusercontent.com/wA5QyOsdZ__aHn7usTK6bpNyHqbUGQWYqQnY1dqi2IdakhGroq86alSgMm_VL-hZREg",
  "banner": "https://play-lh.googleusercontent.com/dqrxZl2HP89EsRcL8Lvgx33GIFeCiVi_GrAYWcR0ewERIg1WdtrkdtoP10UpJv5Gh9WG",
  "privacy_policy": "http://rajkumaar.co.in/privacy.txt",
  "latest_update_message": "No user facing changes in this update, but it contains a mandatory under-the-hood change required by the Google Play Store.",
  "website": "http://rajkumaar.co.in",
  "support_email": "rajkumaar2304@icloud.com"
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

Explore customization options based on shields.io style parameters. Learn more [here](https://shields.io/).

### Examples

Downloads
```markdown
![Downloads Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/downloads?id=in.co.rajkumaar.amritarepo)
```
![Downloads Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/downloads?id=in.co.rajkumaar.amritarepo)

Version
```markdown
![Version Badge](https://img.shields.io/endpoint?color=blue&url=https://api-playstore.rajkumaar.co.in/version?id=in.co.rajkumaar.amritarepo)
```
![Version Badge](https://img.shields.io/endpoint?color=blue&url=https://api-playstore.rajkumaar.co.in/version?id=in.co.rajkumaar.amritarepo)

Developer
```markdown
![Dev Badge](https://img.shields.io/endpoint?color=orange&style=for-the-badge&url=https://api-playstore.rajkumaar.co.in/developer?id=in.co.rajkumaar.amritarepo)
```
![Dev Badge](https://img.shields.io/endpoint?color=orange&style=for-the-badge&url=https://api-playstore.rajkumaar.co.in/developer?id=in.co.rajkumaar.amritarepo)

### License
Find the license [here](LICENSE)



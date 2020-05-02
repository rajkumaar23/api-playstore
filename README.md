# Unofficial PlayStore API

> A helper tool that can be used in combination with shields.io for generating android app related badges with realtime values for your app's README file.

![Playstore API CD](https://github.com/rajkumaar23/playstore-api/workflows/Playstore%20API%20CD/badge.svg)
![Say NO to hardcoding](https://img.shields.io/static/v1?label=Say%20NO%20to&message=HARD-CODING&color=red) 

## Usage

- Sample JSON with all details 
```json
{
  "package": "in.co.rajkumaar.amritarepo",
  "appVersion": "4.2.2",
  "appSize": "4.5M",
  "noOfInstalls": "5,000+",
  "lastUpdated": "March 31, 2020",
  "rating": "4.7",
  "noOfUsersRated": "767",
  "developer": "Rajkumar S"
}
```
To get the above json, you can visit 
```
https://api-playstore.rajkumaar.co.in/json?id=<YOUR-PACKAGE-ID>
```
> Example : https://api-playstore.rajkumaar.co.in/json?id=in.co.rajkumaar.amritarepo

- For a badge on particular attribute, you can use
```markdown
![Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/<ATTRIBUTE-NAME>?id=<PACKAGE-ID>)
```
> The attributes should be one among : [downloads, package, version, size, lastUpdated, rating, noOfUsersRated, developer]

- You can customise the badge as you wish according to the shields.io style parameters. Find more info [here](https://shields.io/)

### Few Examples

- Downloads
```markdown
![Downloads Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/downloads?id=in.co.rajkumaar.amritarepo&color=success)
```
![Downloads Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/downloads?id=in.co.rajkumaar.amritarepo&color=success)

- Version
```markdown
![Version Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/version?id=in.co.rajkumaar.amritarepo&color=blue)
```
![Version Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/version?id=in.co.rajkumaar.amritarepo&color=blue)

- Rating
```markdown
![Rating Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/rating?id=in.co.rajkumaar.amritarepo&color=blueviolet&style=flat-square&logo=android)
```
![Rating Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/rating?id=in.co.rajkumaar.amritarepo&color=blueviolet&style=flat-square&logo=android)

- Developer
```markdown
![Dev Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/developer?id=in.co.rajkumaar.amritarepo&color=orange&style=for-the-badge)
```
![Dev Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/developer?id=in.co.rajkumaar.amritarepo&color=orange&style=for-the-badge)

### License
Find the license [here](LICENSE)



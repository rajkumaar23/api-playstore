# Unofficial PlayStore API

> A helper tool that can be used in combination with shields.io for generating android app related badges with realtime values for your app's README file.

![Custom Badge](https://img.shields.io/static/v1?label=SAY%20NO%20TO&message=HARD-CODING&color=red&style=for-the-badge)

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
https://api-playstore.herokuapp.com/json?id=<YOUR-PACKAGE-ID>
```
> Example : https://api-playstore.herokuapp.com/json?id=in.co.rajkumaar.amritarepo

- For a badge on particular attribute, you can use
```markdown
![Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/<ATTRIBUTE-NAME>?id=<PACKAGE-ID>)
```
> The attributes should be one among : [downloads, package, version, size, lastUpdated, rating, noOfUsersRated, developer]

- You can customise the badge as you wish according to the shields.io style parameters. Find more info [here](https://shields.io/)

### Few Examples

- Downloads
```markdown
![Downloads Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/downloads?id=in.co.rajkumaar.amritarepo&color=success)
```
![Downloads Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/downloads?id=in.co.rajkumaar.amritarepo&color=success)

- Version
```markdown
![Version Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/version?id=in.co.rajkumaar.amritarepo&color=blue)
```
![Version Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/version?id=in.co.rajkumaar.amritarepo&color=blue)

- Rating
```markdown
![Rating Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/rating?id=in.co.rajkumaar.amritarepo&color=blueviolet&style=flat-square&logo=android)
```
![Rating Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/rating?id=in.co.rajkumaar.amritarepo&color=blueviolet&style=flat-square&logo=android)

- Developer
```markdown
![Dev Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/developer?id=in.co.rajkumaar.amritarepo&color=orange&style=for-the-badge)
```
![Dev Badge](https://img.shields.io/endpoint?url=https://api-playstore.herokuapp.com/developer?id=in.co.rajkumaar.amritarepo&color=orange&style=for-the-badge)

### License
Find the license [here](LICENSE)



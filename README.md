# Unofficial PlayStore API

A helper tool that can be used in combination with shields.io for generating android app related badges with realtime values for your app's README file.

![Served hot from my Raspberry Pi](https://img.shields.io/static/v1?label=API%20is%20served%20hot%20from&message=my%20Raspberry%20Pi&color=blue)
![Say NO to hardcoding](https://img.shields.io/static/v1?label=Say%20NO%20to&message=HARD-CODING&color=red)

![Used By Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/used-by)

[Available on RapidAPI](https://rapidapi.com/rajkumaar23-RK1kHf2Zl/api/app-details-from-playstore/details)

## Usage

Sample JSON with all details 
```json
{
  "packageID": "in.co.rajkumaar.amritarepo",
  "developer": "Rajkumar S",
  "installs": "5,000+",
  "lastUpdated": "June 13, 2020",
  "version": "4.3.3"
}
```
To get the above json, you can visit 
```
https://api-playstore.rajkumaar.co.in/json?id=<YOUR-PACKAGE-ID>
```
Example : https://api-playstore.rajkumaar.co.in/json?id=in.co.rajkumaar.amritarepo

For a badge on particular attribute, you can use
```markdown
![Badge](https://img.shields.io/endpoint?url=https://api-playstore.rajkumaar.co.in/<ATTRIBUTE-NAME>?id=<PACKAGE-ID>)
```
The attributes should be one among : [downloads, package, version, lastUpdated, developer]

You can customise the badge as you wish according to the shields.io style parameters. Find more info [here](https://shields.io/)

### Few Examples

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



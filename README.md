# Unofficial PlayStore API

A helper tool that can be used in combination with shields.io for generating android app related badges with realtime values for your app's README file.

[![Netlify Status](https://api.netlify.com/api/v1/badges/24ba080b-cb3b-42f5-ba62-97c546333aa8/deploy-status)](https://app.netlify.com/sites/practical-mcclintock-bfa26d/deploys)
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

### API Used by
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/rajkumaar23/amrita-repository"><img src="https://raw.githubusercontent.com/rajkumaar23/amrita-repository/master/app/src/main/res/drawable/logosq.png" width="100px;" alt=""/><br/><sub><b>Amrita Repository</b></sub></a></td>
    <td align="center"><a href="https://github.com/capturemathan/PocketFeed"><img src="https://raw.githubusercontent.com/capturemathan/PocketFeed/master/app/src/main/res/mipmap-xxxhdpi/ic_launcher_foreground.png" width="100px;" alt=""/><br/><sub><b>PocketFeed</b></sub></a></td>
    <td align="center"><a href="https://github.com/leoxshn/posidonLauncher"><img src="https://raw.githubusercontent.com/leoxshn/posidonLauncher/master/fastlane/metadata/android/en-US/images/icon.png" width="100px;" alt=""/><br/><sub><b>posidon launcher</b></sub></a></td>
  </tr>
</table>
<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->

Submit a PR with your app details if you are using this API.

### License
Find the license [here](LICENSE)



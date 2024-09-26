# Unofficial PlayStore API

**Explore android app metadata from PlayStore with ease!**

[![Build & Deploy](https://github.com/rajkumaar23/api-playstore/actions/workflows/deploy.yaml/badge.svg)](https://github.com/rajkumaar23/api-playstore/actions/workflows/deploy.yaml)
[![Nightly Test](https://github.com/rajkumaar23/api-playstore/actions/workflows/test.yaml/badge.svg)](https://github.com/rajkumaar23/api-playstore/actions/workflows/test.yaml)

## Usage

Retrieve comprehensive JSON data for any app from Google Play by making a simple GET request:
```
https://api-playstore.rajkumaar.co.in/json?id=<YOUR-PACKAGE-ID>
```
Example : https://api-playstore.rajkumaar.co.in/json?id=com.dd.doordash

Sample JSON response with all possible details:
```json
{
    "packageID": "com.dd.doordash",
    "name": "DoorDash - Food Delivery",
    "version": "15.156.24",
    "downloads": "50,000,000+",
    "downloadsExact": 61535856,
    "lastUpdated": "Mar 29, 2024",
    "launchDate": "Mar 26, 2015",
    "developer": "DoorDash",
    "screenshots": [
        "https://play-lh.googleusercontent.com/b-ZRIvKPErodCjPJo47AB_Fl1dGgOUAGoe03OjER7lzZPffz3Cv0xgMniNHfRYR-bQg",
        "https://play-lh.googleusercontent.com/-z9S93iM5RR6xOU8Sj774aNZsSATCZymn3fJiOnzW1VuGteh5vLfgyTjIiKOMhCFfjc",
        "https://play-lh.googleusercontent.com/2iNMFwIwCHD8otkfT_AgZU62UNqJziCay3uEVEDFrXmz78wuOf-hTNZbLwdEgShWcfM",
        "https://play-lh.googleusercontent.com/UZPWjpQcp-3U2FeUxAIuU9zX2P-z4rhKwmH0hyyr4sUQoKVUotKb_SdhuRBxMW0KsuY",
        "https://play-lh.googleusercontent.com/jUcGpu7azj7RwB4brZyq1xoCzyobyUkHwyNcuhfZJK-6DDvm7AFB0ldsTP4Hs2h5UsHg",
        "https://play-lh.googleusercontent.com/3Fkc_6o-818jRxVuP42rluJNjx_CQ77Bs8Fx-HDD6jAxRmD-irJLsMAtiIvm3sC8qeZN"
    ],
    "category": "Food \u0026 Drink",
    "logo": "https://play-lh.googleusercontent.com/Fvled-zLfL8ER0EBNIk-FnunJCcH2u_T6rdITclOFdU2jpEopdMstZOHP-PtuhP_5coZ",
    "banner": "https://play-lh.googleusercontent.com/TmWukTjL87iibmuHXhc9d04Hvul6nErQXE0nRcJvvsiFffDBjE0JmGns57r1x1RihMA",
    "privacy_policy": "https://www.doordash.com/privacy/",
    "latest_update_message": "Our mission is to deliver good by connecting people and possibility. If we can play a small role in helping you spend more time with your friends and family or get ahead on your favorite projects, then we have delivered good.\u003cbr\u003e\u003cbr\u003eWe’re launching a new set of initiatives to deliver good within our communities, and you’ll see a new visual identity that reflects our spirit. Thank you for making DoorDash a success — we wouldn’t be here without you.\u003cbr\u003e\u003cbr\u003eUpward and onward!\u003cbr\u003eTony Xu, CEO and Co-Founder",
    "website": "https://www.doordash.com/",
    "support_email": "support@doordash.com",
    "rating": "4.6",
    "noOfUsersRated": "4,299,422",
    "description": "Delivery anywhere you are. DoorDash offers the greatest online selection of your favorite restaurants and stores, facilitating delivery of freshly prepared meals, groceries, OTC medicines, flowers \u0026amp; more. With more than 310,000 menus and 55,000+ grocery, convenience \u0026amp; retail stores across 4,000+ cities in the U.S., Canada, and Australia, you’ll find the best of your neighborhood as you shop and order online. Plus, enjoy $0 delivery fees for your first month. Restrictions apply: https://drd.sh/tF5uns/\u003cbr\u003e\u003cbr\u003eIT’S ALL HERE\u003cbr\u003e·Restaurants: Food delivery from local \u0026amp; national restaurants\u003cbr\u003e·Grocery: Fulfill your weekly grocery list, from produce to diapers. Shop diets like gluten-free \u0026amp; vegan.\u003cbr\u003e·Drinks, Snacks \u0026amp; More: Sodas, candy \u0026amp; ibuprofen from stores like 7-Eleven \u0026amp; CVS\u003cbr\u003e·Flowers: Order fresh flowers for occasions like Valentine’s Day\u003cbr\u003e\u003cbr\u003eKEY FEATURES\u003cbr\u003e· Get it now: Get on-demand, same-day delivery\u003cbr\u003e· Schedule deliveries: Order in advance and get it when it\u0026#39;s most convenient for you.\u003cbr\u003e· Real-Time Tracking: See when your order will arrive.\u003cbr\u003e· No Minimums: Order as little or as much as you want.\u003cbr\u003e· Easy Payment: Conveniently pay via Google Pay, Venmo, Paypal, credit card, or SNAP/EBT at participating Mx.\u003cbr\u003e· No-Contact Delivery: Your food will be left in a safe place and you’ll be alerted that it’s ready for you to pick up at your doorstep\u003cbr\u003e\u003cbr\u003eENJOY UNLIMITED $0 DELIVERY FEES WITH DASHPASS\u003cbr\u003eGet unlimited $0 delivery fees and up to 10% off eligible orders from your neighborhood restaurants, grocery stores, and more. Plus, DashPass members get access to exclusive items and offers, and 5% in DoorDash credits back on eligible Pickup orders. Your first 30 days on DashPass are free, then your membership auto-renews at $9.99/month. Cancel anytime.\u003cbr\u003e\u003cbr\u003eNATIONAL RESTAURANT PARTNERS\u003cbr\u003eMcDonald\u0026#39;s, Starbucks, Chick-fil-A, Burger King, Wendy’s, Chipotle, The Cheesecake Factory, Outback Steakhouse, Panera, Chili\u0026#39;s, Subway, Dunkin’ Donuts, Jamba Juice, Panda Express, Moe\u0026#39;s, P.F. Chang’s, Denny’s, Buffalo Wild Wings, Papa John\u0026#39;s, Papa Murphy’s, Jack in the Box, Five Guys, Boston Market, Red Robin, TGI Friday’s, Red Lobster, Qdoba, El Pollo Loco, White Castle, SmashBurger\u003cbr\u003e\u003cbr\u003eGROCERY DELIVERY PARTNERS\u003cbr\u003eSafeway, Albertsons, Aldi, Sprouts Farmers Market, Meijer, Hy-Vee, Grocery Outlet, Winn-Dixie, Smart \u0026amp; Final, BJ’s, Vons, Weis, ACME, Raley’s, Fresh Thyme, Giant Eagle, Bashas’, Bristol Farms and more.\u003cbr\u003e\u003cbr\u003eCONVENIENCE \u0026amp; RETAIL DELIVERY PARTNERS\u003cbr\u003eWalgreens, 7-Eleven, CVS, Rite Aid, Dollar General, Wawa, Sheetz, Casey’s, Total Wine, BevMo!, PetSmart, Sephora, DICK\u0026#39;S Sporting Goods, Tractor Supply, and more\u003cbr\u003e\u003cbr\u003eFIND RESTAURANTS AND STORES NEAR YOU\u003cbr\u003eWe’re growing and currently serving over 4,000 cities across the United States, Puerto Rico, Canada, and Australia including cities such as New York City, Los Angeles, Toronto, Vancouver, BC, Melbourne, Sydney, Montreal and more.\u003cbr\u003e\u003cbr\u003eNotice at Collection (California Residents): https://help.doordash.com/consumers/s/privacy-policy-us#section-11\u003cbr\u003eVisit doordash.com to learn more.",
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
![Downloads Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/downloads?id=com.dd.doordash)
```
![Downloads Badge](https://img.shields.io/endpoint?color=success&url=https://api-playstore.rajkumaar.co.in/downloads?id=com.dd.doordash)

Version
```markdown
![Version Badge](https://img.shields.io/endpoint?color=blue&url=https://api-playstore.rajkumaar.co.in/version?id=com.dd.doordash)
```
![Version Badge](https://img.shields.io/endpoint?color=blue&url=https://api-playstore.rajkumaar.co.in/version?id=com.dd.doordash)

Developer
```markdown
![Dev Badge](https://img.shields.io/endpoint?color=orange&style=for-the-badge&url=https://api-playstore.rajkumaar.co.in/developer?id=com.dd.doordash)
```
![Dev Badge](https://img.shields.io/endpoint?color=orange&style=for-the-badge&url=https://api-playstore.rajkumaar.co.in/developer?id=com.dd.doordash)

### Known Issues

- The API may respond with a **0** for `rating`/`noOfUsersRated` because the ratings on Google Play Store are specific to a country/region and the API is only able to fetch ratings that are local to the United States.

### License
Find the license [here](LICENSE)



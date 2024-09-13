package credly_test

import (
	"testing"

	"github.com/mikejoh/go-credly/internal/credly"
)

func TestExtractBadges(t *testing.T) {
	tt := []struct {
		name     string
		html     string
		expected []credly.Badge
	}{
		{
			name: "extract badges",
			html: credlyHTML,
			expected: []credly.Badge{
				{
					ImageSrc: "https://images.credly.com/size/110x110/images/8b8ed108-e77d-4396-ac59-2504583b9d54/cka_from_cncfsite__281_29.png",
					Alt:      "",
				},
				{
					ImageSrc: "https://images.credly.com/size/110x110/images/f28f1d88-428a-47f6-95b5-7da1dd6c1000/KCNA_badge.png",
					Alt:      "",
				},
			},
		},

		{
			name:     "no badges",
			html:     "",
			expected: nil,
		},

		{
			name:     "invalid HTML",
			html:     "<html></html>",
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			badges, err := credly.ExtractBadges([]byte(tc.html))
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if len(badges) != len(tc.expected) {
				t.Fatalf("expected %d badges, got %d", len(tc.expected), len(badges))
			}

			for i, badge := range badges {
				if badge.ImageSrc != tc.expected[i].ImageSrc {
					t.Fatalf("expected image src %s, got %s", tc.expected[i].ImageSrc, badge.ImageSrc)
				}

				if badge.Alt != tc.expected[i].Alt {
					t.Fatalf("expected alt %s, got %s", tc.expected[i].Alt, badge.Alt)
				}
			}
		})
	}
}

// 2024-09-13
var credlyHTML = `<!DOCTYPE html>
<html lang='en'>
<head>
<meta charset='UTF-8'>
  <script src="https://cdn.cookielaw.org/scripttemplates/otSDKStub.js"  type="text/javascript" charset="UTF-8" data-domain-script="47d1c5ce-460c-434d-8cc7-0318d7d02a9d" ></script>
  <script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[

    function OptanonWrapper() { }

//]]>
</script>
<script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[

        window._trackJs = { token: 'e43bb3a0b11a469ba1ef2ec2b79c6b11', application: 'acclaim-production', version: '1.90.2'};

//]]>
</script><script type="text/javascript" src="https://cdn.trackjs.com/releases/current/tracker.js"></script>


<script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[

    (function(apiKey){
      (function(p,e,n,d,o){var v,w,x,y,z;o=p[d]=p[d]||{};o._q=[];
        v=['initialize','identify','updateOptions','pageLoad'];for(w=0,x=v.length;w < x;++w)(function(m){
        o[m]=o[m]||function(){o._q[m===v[0]?'unshift':'push']([m].concat([].slice.call(arguments,0)));};})(v[w]);
        y=e.createElement(n);y.async=!0;y.src='https://cdn.pendo.io/agent/static/'+apiKey+'/pendo.js';
        z=e.getElementsByTagName(n)[0];z.parentNode.insertBefore(y,z);})(window,document,'script','pendo');
    })("65460ef3-56d0-45a3-7b3b-fe1ec0463054")

//]]>
</script>
<meta content='Credly is a global Open Badge platform that closes the gap between skills and opportunities. We work with academic institutions, corporations, and professional associations to translate learning outcomes into digital credentials that are immediately validated, managed, and shared.' data-rh name='description'>
<meta content='Credly, Credly badges, Acclaim, open badges, digital badges, badges, web-enabled credentials, Badge Alliance, Mozilla Open Badge standards, Mozilla Open Badges, verified credentials, learning outcomes' name='keywords'>
<meta content='width=device-width' name='viewport'>
<meta content='1B97012E4CB3B07611090D1A0B4D9D19' name='msvalidate.01'>
<meta content='d20lZjrSJJ_n0jc1HHlADZBDmn5wZfiBDucFzzOzHCY' name='google-site-verification'>
<meta property="og:title" content="Mikael Johansson&#39;s profile on Credly">
<meta property="og:image" content="https://cdn.credly.com/assets/empties/user-903c91badbde74c6f6354116d1d2ad45051ad281b16e04bf7bc2a3c41653af4c.png">
<meta property="og:description" content="View Mikael Johansson&#39;s verified achievements on Credly. ">
<meta property="profile:first_name" content="Mikael">
<meta property="profile:last_name" content="Johansson">
<meta property="og:type" content="og:profile">
<meta property="og:url" content="https://www.credly.com/users/a52c3dc9-d0b3-4f23-b319-14756ed00dd0">
<meta property="og:site_name" content="Credly">
<meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="zZeEpSupCKUTF6aRI2vp4sfnJ-vd-Irc913FMRMvBT7c-q3JmQqemKr2sbRq9EqqQoHECcq_VGr4uwAvxWnyPw" />
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Open+Sans:300,400,400i,600,700" media="all" />
<link rel="stylesheet" href="//use.typekit.net/bcc0eop.css" media="all" />
<script async='true' src='https://www.googletagmanager.com/gtag/js?id=G-5D8QQR0C2M'></script>
<script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[
window.dataLayer = window.dataLayer || [];
function google_tag(){dataLayer.push(arguments);}
google_tag('js', new Date());
google_tag('config', 'G-5D8QQR0C2M');
(function() {
var tests = [];
for (var k in {}) {tests.push(k);}
google_tag('set', 'dimension1', tests.join(';'));
})();
google_tag('send', 'pageview');

//]]>
</script><title>Mikael Johansson - Credly</title>
<link rel="stylesheet" href="https://cdn.credly.com/assets/transitional-77168a2fe1d4bebcc860ef4ecbe16637defc7c27710e5ffb8dda025fa71c1c69.css" media="all" />
<link rel="stylesheet" href="https://cdn.credly.com/assets/application-a9882c483022f1c09157ef3f2b986df428d5f183fb0129f20e308e0b463767c7.css" media="all" />
</head>
<body class=''>
<script src="https://cdn.credly.com/assets/ie_polyfills.legacy-browser-19f8e03888197c91592c12b27330a2469145e9c2704aacb53a996c27be0e046c.js" crossorigin="anonymous"></script>
<script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[
setTimeout(function(){pendo.initialize({"visitor":{},"account":{}});}, 1000)

//]]>
</script><header class='transitional' id='website-header' style=''>
<div class='website-header-main-links-container website-header-placeholder'></div>
<div>
<a href='#skip-target' id='skip-to-content'>Skip to content</a>
</div>
<div class='website-header-main-links-container non-react-header'>
<div class='grid website-header-main-links'>
<a class="acclaim-logo" href="https://www.credly.com/">Credly</a>
<nav data-behavior='container' name='main_navigation'>
<div class='header-nav-item header-nav-item--btn-container header-nav-item--btn-container--create-account-btn-container'>
<a class='create-account-btn button white' href='/users/sign_up'>
Create Account
</a>

</div>
<div class='header-nav-item header-nav-item--btn-container header-nav-item--btn-container--sign-in-btn-container'>
<a class='sign-in-btn button' href='/users/sign_in'>
Sign In
</a>

</div>
</nav>
</div>
</div>

</header>
<div class='search-bar' data-container='search-bar'></div>
<div class='user_public_badges' id='flash'>
<div class='grid'>
<div class='l1-0 r1-0' data-behavior='flash-container'>
<noscript>
<div class='flash notice'>
<i class='icon-error'></i>
<div>
We've detected that your browser has JavaScript disabled. Some features of Credly require Javascript to be enabled. For the best experience, please enable JavaScript in your browser settings or try using a different browser.
</div>
</div>
</noscript>

</div>
</div>
</div>
<main id="root">

<div class='container'>
<div class='c-profile-header'>
<div class='c-profile-header__image-wrap cr-avatar__image-wrap'>
<img alt="Mikael Johansson" class="cr-avatar__image--auto" src="https://cdn.credly.com/assets/empties/user_new-68d5833b1b59b8724c643893b1dccfe9ca1b652286ccf964ff211a4a820e2b08.svg" />
</div>
<div class='c-profile-header__detail'>
<h1 class="ac-heading ac-heading--badge-name-hero">Mikael Johansson
</h1><div class='c-profile-header__subtitles'>
<div class='c-profile-header__subtitle'>
Site Reliability Engineer | Etraveli Group
</div>
<div class='c-profile-header__subtitle'>
Gothenburg
</div>
</div>
<div class='c-profile-header__description'>

</div>
<div class='c-profile-header__links'>





</div>
</div>
</div>


</div>
<nav class='c-top-nav'>
<ul>
<li class='c-top-nav__link c-top-nav__link--active'>
<a class="c-top-nav__link c-top-nav__link--active" href="/users/mikael-johansson-2/badges">Badges
</a></li>
</ul>
</nav>

<div class='container'>
<div class='row justify-content-between align-items-center cr-public-list-view__count-row'>
<div class='col-auto'>
2 badges
</div>
<div class="c-selector-menu col-auto"><div class="c-selector-menu__label" for="selector_menu_c766"><span>Sort by:&nbsp;</span></div><div><span class="c-selector-menu__current-selection" aria-labeledby="selector_menu_c766" tabindex="0">Most popular<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><path d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"/></svg></span><ul class="c-selector-menu__options"><li class="c-selector-menu__option c-selector-menu__option--selected"><a href="/users/mikael-johansson-2/badges?id=mikael-johansson-2&amp;page=1&amp;sort=most_popular">Most popular</a></li><li class="c-selector-menu__option"><a href="/users/mikael-johansson-2/badges?id=mikael-johansson-2&amp;page=1&amp;sort=-state_updated_at">Most recent</a></li></ul></div></div>
</div>
<div class="data-table data-table-grid"><ul class="data-table__rows data-table-grid__rows"><li class="data-table-row data-table-row-grid"><div class="col-12 col data-table-content"><a class="c-grid-item c-grid-item--stack-lt-sm cr-public-earned-badge-grid-item" title="CKA: Certified Kubernetes Administrator" href="/badges/20f4aaea-770e-4e32-8cd0-f2720fb11d85"><div class='cr-standard-grid-item-content c-badge c-badge--medium'>
<img class="cr-standard-grid-item-content__image" alt="" src="https://images.credly.com/size/110x110/images/8b8ed108-e77d-4396-ac59-2504583b9d54/cka_from_cncfsite__281_29.png" />
<div class='cr-standard-grid-item-content__details'>
<div class='cr-standard-grid-item-content__title'>
CKA: Certified Kubernetes Administrator
</div>
<div class='cr-standard-grid-item-content__subtitle'>
The Linux Foundation
</div>
</div>
</div>


</a>
</div></li><li class="data-table-row data-table-row-grid"><div class="col-12 col data-table-content"><a class="c-grid-item c-grid-item--stack-lt-sm cr-public-earned-badge-grid-item" title="KCNA: Kubernetes and Cloud Native Associate" href="/badges/062ae104-f532-43d0-b3bd-b6599dd03e2c"><div class='cr-standard-grid-item-content c-badge c-badge--medium'>
<img class="cr-standard-grid-item-content__image" alt="" src="https://images.credly.com/size/110x110/images/f28f1d88-428a-47f6-95b5-7da1dd6c1000/KCNA_badge.png" />
<div class='cr-standard-grid-item-content__details'>
<div class='cr-standard-grid-item-content__title'>
KCNA: Kubernetes and Cloud Native Associate
</div>
<div class='cr-standard-grid-item-content__subtitle'>
The Linux Foundation
</div>
</div>
</div>


</a>
</div></li></ul></div><div class='row'>

</div>


</div>

</main><div data-container='global-search-results'></div>
<footer class='transitional' id='footer'>
<div class='grid non-react-footer'>
<section class='l1-0 r1-0 footer-margin'>
<div class='links'>
<div class='footer-link-wrapper'>
<a class="footer-link" href="https://resources.credly.com/schedule-a-demo">Request Demo</a>
</div>
<div class='footer-link-wrapper'>
<a class="footer-link" href="https://info.credly.com/about-us">About Credly</a>
</div>
<div class='footer-link-wrapper'>
<a class="footer-link" href="https://info.credly.com/legal">Terms</a>
</div>
<div class='footer-link-wrapper'>
<a class="footer-link" href="https://info.credly.com/privacy-policy">Privacy</a>
</div>
<div class='footer-link-wrapper'>
<a class="footer-link" href="/docs">Developers</a>
</div>
<div class='footer-link-wrapper'>
<a class="footer-link" href="/support">Support</a>
</div>
</div>
<div class='cookie-links'>
<div class='cookie-link-wrapper'>
<a class="footer-link" target="_blank" href="https://info.credly.com/cookie-policy">Cookies</a>
</div>
<div class='cookie-link-wrapper'>
<a class="footer-link" target="_blank" href="https://info.credly.com/cookie-policy">Do Not Sell My Personal Information</a>
</div>
</div>
<div class='app-store-links'>
<a class='app-store-link app-store-link-google' href='https://play.google.com/store/apps/details?id=com.credly.android'></a>
<a class='app-store-link app-store-link-apple' href='https://apps.apple.com/us/app/credly/id1630234370'></a>
</div>
</section>
</div>

<span></span>
</footer>
<script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[
(function(){
  document.addEventListener('click', function(e){
    var eventTarget = e.target;
    if (eventTarget.id == 'skip-to-content' || eventTarget.id == 'skip-nav'){
      var href, target;
      href = eventTarget.getAttribute('href').substr(1);
      target = document.getElementById(href);
      if (!target.getAttribute('tabindex')){
        target.setAttribute('tabindex', '-1');
      }
      target.focus();
    }
  });
})();


//]]>
</script><script src="https://cdn.credly.com/assets/utilities/set_time_zone_cookie-aca1fcb79ea4e90d8e963d7850b867a9bbeb077b05a6cad141c1ce7d945fd20d.js" crossorigin="anonymous"></script>
<script nonce="+YSX/2gcigYfe4WKRc8SRS/yM1jhERP90EuvBXQ4Jlc=">
//<![CDATA[
(function(){
  if (typeof(App) !== 'undefined' && App.Behaviors && App.Behaviors.TrackStat) {
    App.Behaviors.TrackStat.init(
      {"url":"https:\/\/stats.credly.com\/stats\/interaction","request_data":{"session_id":"023f4576-0f1e-4c5e-a47b-3fe75a29d783","flexible_params_auth":["snapshot_json","stat_object_id","stat_object_type","stat_type"],"auth_version":"1.0","auth_key":"user:8d768abb-4e6a-4f38-8cb3-86dc95196f75","auth_timestamp":"1726236354","auth_signature":"45fed5237f097e9c1f19932662923978ecd6dba3d3e014a50eb0e91353fd283a"}},
        null
    );
  }
})();


(function(){
  const vals = window.initialAppValues = window.initialAppValues || {};
  vals.locationData = {
    allowedCountries: JSON.parse("[{\"id\":227,\"name\":\"United States\",\"zip_required\":false,\"org_state_required\":true},{\"id\":38,\"name\":\"Canada\",\"zip_required\":false,\"org_state_required\":true},{\"id\":3,\"name\":\"Afghanistan\",\"zip_required\":false},{\"id\":15,\"name\":\"Aland Islands\",\"zip_required\":false},{\"id\":6,\"name\":\"Albania\",\"zip_required\":false},{\"id\":61,\"name\":\"Algeria\",\"zip_required\":false},{\"id\":11,\"name\":\"American Samoa\",\"zip_required\":false},{\"id\":1,\"name\":\"Andorra\",\"zip_required\":false},{\"id\":8,\"name\":\"Angola\",\"zip_required\":false},{\"id\":5,\"name\":\"Anguilla\",\"zip_required\":false},{\"id\":9,\"name\":\"Antarctica\",\"zip_required\":false},{\"id\":4,\"name\":\"Antigua and Barbuda\",\"zip_required\":false},{\"id\":10,\"name\":\"Argentina\",\"zip_required\":false},{\"id\":7,\"name\":\"Armenia\",\"zip_required\":false},{\"id\":14,\"name\":\"Aruba\",\"zip_required\":false},{\"id\":13,\"name\":\"Australia\",\"zip_required\":false},{\"id\":12,\"name\":\"Austria\",\"zip_required\":false},{\"id\":16,\"name\":\"Azerbaijan\",\"zip_required\":false},{\"id\":32,\"name\":\"Bahamas\",\"zip_required\":false},{\"id\":23,\"name\":\"Bahrain\",\"zip_required\":false},{\"id\":19,\"name\":\"Bangladesh\",\"zip_required\":false},{\"id\":18,\"name\":\"Barbados\",\"zip_required\":false},{\"id\":36,\"name\":\"Belarus\",\"zip_required\":false},{\"id\":20,\"name\":\"Belgium\",\"zip_required\":false},{\"id\":37,\"name\":\"Belize\",\"zip_required\":false},{\"id\":25,\"name\":\"Benin\",\"zip_required\":false},{\"id\":27,\"name\":\"Bermuda\",\"zip_required\":false},{\"id\":33,\"name\":\"Bhutan\",\"zip_required\":false},{\"id\":29,\"name\":\"Bolivia\",\"zip_required\":false},{\"id\":30,\"name\":\"Bonaire, Sint Eustatius and Saba\",\"zip_required\":false},{\"id\":17,\"name\":\"Bosnia and Herzegovina\",\"zip_required\":false},{\"id\":35,\"name\":\"Botswana\",\"zip_required\":false},{\"id\":34,\"name\":\"Bouvet Island\",\"zip_required\":false},{\"id\":31,\"name\":\"Brazil\",\"zip_required\":false},{\"id\":105,\"name\":\"British Indian Ocean Territory\",\"zip_required\":false},{\"id\":28,\"name\":\"Brunei Darussalam\",\"zip_required\":false},{\"id\":22,\"name\":\"Bulgaria\",\"zip_required\":false},{\"id\":21,\"name\":\"Burkina Faso\",\"zip_required\":false},{\"id\":24,\"name\":\"Burundi\",\"zip_required\":false},{\"id\":115,\"name\":\"Cambodia\",\"zip_required\":false},{\"id\":47,\"name\":\"Cameroon\",\"zip_required\":false},{\"id\":51,\"name\":\"Cape Verde\",\"zip_required\":false},{\"id\":121,\"name\":\"Cayman Islands\",\"zip_required\":false},{\"id\":41,\"name\":\"Central African Republic\",\"zip_required\":false},{\"id\":209,\"name\":\"Chad\",\"zip_required\":false},{\"id\":46,\"name\":\"Chile\",\"zip_required\":false},{\"id\":48,\"name\":\"China\",\"zip_required\":false},{\"id\":53,\"name\":\"Christmas Island\",\"zip_required\":false},{\"id\":39,\"name\":\"Cocos (Keeling) Islands\",\"zip_required\":false},{\"id\":49,\"name\":\"Colombia\",\"zip_required\":false},{\"id\":117,\"name\":\"Comoros\",\"zip_required\":false},{\"id\":42,\"name\":\"Congo\",\"zip_required\":false},{\"id\":40,\"name\":\"Congo, The Democratic Republic Of The\",\"zip_required\":false},{\"id\":45,\"name\":\"Cook Islands\",\"zip_required\":false},{\"id\":50,\"name\":\"Costa Rica\",\"zip_required\":false},{\"id\":44,\"name\":\"Côte D'Ivoire\",\"zip_required\":false},{\"id\":97,\"name\":\"Croatia\",\"zip_required\":false},{\"id\":52,\"name\":\"Curaçao\",\"zip_required\":false},{\"id\":54,\"name\":\"Cyprus\",\"zip_required\":false},{\"id\":55,\"name\":\"Czech Republic\",\"zip_required\":false},{\"id\":58,\"name\":\"Denmark\",\"zip_required\":false},{\"id\":57,\"name\":\"Djibouti\",\"zip_required\":false},{\"id\":59,\"name\":\"Dominica\",\"zip_required\":false},{\"id\":60,\"name\":\"Dominican Republic\",\"zip_required\":false},{\"id\":62,\"name\":\"Ecuador\",\"zip_required\":false},{\"id\":64,\"name\":\"Egypt\",\"zip_required\":false},{\"id\":205,\"name\":\"El Salvador\",\"zip_required\":false},{\"id\":87,\"name\":\"Equatorial Guinea\",\"zip_required\":false},{\"id\":66,\"name\":\"Eritrea\",\"zip_required\":false},{\"id\":63,\"name\":\"Estonia\",\"zip_required\":false},{\"id\":68,\"name\":\"Ethiopia\",\"zip_required\":false},{\"id\":71,\"name\":\"Falkland Islands (Malvinas)\",\"zip_required\":false},{\"id\":73,\"name\":\"Faroe Islands\",\"zip_required\":false},{\"id\":70,\"name\":\"Fiji\",\"zip_required\":false},{\"id\":69,\"name\":\"Finland\",\"zip_required\":false},{\"id\":74,\"name\":\"France\",\"zip_required\":false},{\"id\":79,\"name\":\"French Guiana\",\"zip_required\":false},{\"id\":172,\"name\":\"French Polynesia\",\"zip_required\":false},{\"id\":210,\"name\":\"French Southern Territories\",\"zip_required\":false},{\"id\":75,\"name\":\"Gabon\",\"zip_required\":false},{\"id\":84,\"name\":\"Gambia\",\"zip_required\":false},{\"id\":78,\"name\":\"Georgia\",\"zip_required\":false},{\"id\":56,\"name\":\"Germany\",\"zip_required\":false},{\"id\":81,\"name\":\"Ghana\",\"zip_required\":false},{\"id\":82,\"name\":\"Gibraltar\",\"zip_required\":false},{\"id\":88,\"name\":\"Greece\",\"zip_required\":false},{\"id\":83,\"name\":\"Greenland\",\"zip_required\":false},{\"id\":77,\"name\":\"Grenada\",\"zip_required\":false},{\"id\":86,\"name\":\"Guadeloupe\",\"zip_required\":false},{\"id\":91,\"name\":\"Guam\",\"zip_required\":false},{\"id\":90,\"name\":\"Guatemala\",\"zip_required\":false},{\"id\":80,\"name\":\"Guernsey\",\"zip_required\":false},{\"id\":85,\"name\":\"Guinea\",\"zip_required\":false},{\"id\":92,\"name\":\"Guinea-Bissau\",\"zip_required\":false},{\"id\":93,\"name\":\"Guyana\",\"zip_required\":false},{\"id\":98,\"name\":\"Haiti\",\"zip_required\":false},{\"id\":95,\"name\":\"Heard and McDonald Islands\",\"zip_required\":false},{\"id\":230,\"name\":\"Holy See (Vatican City State)\",\"zip_required\":false},{\"id\":96,\"name\":\"Honduras\",\"zip_required\":false},{\"id\":94,\"name\":\"Hong Kong\",\"zip_required\":false},{\"id\":99,\"name\":\"Hungary\",\"zip_required\":false},{\"id\":107,\"name\":\"Iceland\",\"zip_required\":false},{\"id\":104,\"name\":\"India\",\"zip_required\":false},{\"id\":100,\"name\":\"Indonesia\",\"zip_required\":false},{\"id\":106,\"name\":\"Iraq\",\"zip_required\":false},{\"id\":101,\"name\":\"Ireland\",\"zip_required\":false},{\"id\":103,\"name\":\"Isle of Man\",\"zip_required\":false},{\"id\":102,\"name\":\"Israel\",\"zip_required\":false},{\"id\":108,\"name\":\"Italy\",\"zip_required\":false},{\"id\":110,\"name\":\"Jamaica\",\"zip_required\":false},{\"id\":112,\"name\":\"Japan\",\"zip_required\":false},{\"id\":109,\"name\":\"Jersey\",\"zip_required\":false},{\"id\":111,\"name\":\"Jordan\",\"zip_required\":false},{\"id\":122,\"name\":\"Kazakhstan\",\"zip_required\":false},{\"id\":113,\"name\":\"Kenya\",\"zip_required\":false},{\"id\":116,\"name\":\"Kiribati\",\"zip_required\":false},{\"id\":119,\"name\":\"Korea, Republic of\",\"zip_required\":false},{\"id\":283,\"name\":\"Kosovo\",\"zip_required\":false},{\"id\":120,\"name\":\"Kuwait\",\"zip_required\":false},{\"id\":114,\"name\":\"Kyrgyzstan\",\"zip_required\":false},{\"id\":123,\"name\":\"Lao People's Democratic Republic\",\"zip_required\":false},{\"id\":132,\"name\":\"Latvia\",\"zip_required\":false},{\"id\":124,\"name\":\"Lebanon\",\"zip_required\":false},{\"id\":129,\"name\":\"Lesotho\",\"zip_required\":false},{\"id\":128,\"name\":\"Liberia\",\"zip_required\":false},{\"id\":133,\"name\":\"Libya\",\"zip_required\":false},{\"id\":126,\"name\":\"Liechtenstein\",\"zip_required\":false},{\"id\":130,\"name\":\"Lithuania\",\"zip_required\":false},{\"id\":131,\"name\":\"Luxembourg\",\"zip_required\":false},{\"id\":145,\"name\":\"Macao\",\"zip_required\":false},{\"id\":139,\"name\":\"Madagascar\",\"zip_required\":false},{\"id\":153,\"name\":\"Malawi\",\"zip_required\":false},{\"id\":155,\"name\":\"Malaysia\",\"zip_required\":false},{\"id\":152,\"name\":\"Maldives\",\"zip_required\":false},{\"id\":142,\"name\":\"Mali\",\"zip_required\":false},{\"id\":150,\"name\":\"Malta\",\"zip_required\":false},{\"id\":140,\"name\":\"Marshall Islands\",\"zip_required\":false},{\"id\":147,\"name\":\"Martinique\",\"zip_required\":false},{\"id\":148,\"name\":\"Mauritania\",\"zip_required\":false},{\"id\":151,\"name\":\"Mauritius\",\"zip_required\":false},{\"id\":240,\"name\":\"Mayotte\",\"zip_required\":false},{\"id\":154,\"name\":\"Mexico\",\"zip_required\":false},{\"id\":72,\"name\":\"Micronesia, Federated States Of\",\"zip_required\":false},{\"id\":136,\"name\":\"Moldova, Republic of\",\"zip_required\":false},{\"id\":135,\"name\":\"Monaco\",\"zip_required\":false},{\"id\":144,\"name\":\"Mongolia\",\"zip_required\":false},{\"id\":137,\"name\":\"Montenegro\",\"zip_required\":false},{\"id\":149,\"name\":\"Montserrat\",\"zip_required\":false},{\"id\":134,\"name\":\"Morocco\",\"zip_required\":false},{\"id\":156,\"name\":\"Mozambique\",\"zip_required\":false},{\"id\":143,\"name\":\"Myanmar\",\"zip_required\":false},{\"id\":157,\"name\":\"Namibia\",\"zip_required\":false},{\"id\":166,\"name\":\"Nauru\",\"zip_required\":false},{\"id\":165,\"name\":\"Nepal\",\"zip_required\":false},{\"id\":163,\"name\":\"Netherlands\",\"zip_required\":false},{\"id\":158,\"name\":\"New Caledonia\",\"zip_required\":false},{\"id\":168,\"name\":\"New Zealand\",\"zip_required\":false},{\"id\":162,\"name\":\"Nicaragua\",\"zip_required\":false},{\"id\":159,\"name\":\"Niger\",\"zip_required\":false},{\"id\":161,\"name\":\"Nigeria\",\"zip_required\":false},{\"id\":167,\"name\":\"Niue\",\"zip_required\":false},{\"id\":160,\"name\":\"Norfolk Island\",\"zip_required\":false},{\"id\":146,\"name\":\"Northern Mariana Islands\",\"zip_required\":false},{\"id\":141,\"name\":\"North Macedonia\",\"zip_required\":false},{\"id\":164,\"name\":\"Norway\",\"zip_required\":false},{\"id\":169,\"name\":\"Oman\",\"zip_required\":false},{\"id\":175,\"name\":\"Pakistan\",\"zip_required\":false},{\"id\":182,\"name\":\"Palau\",\"zip_required\":false},{\"id\":180,\"name\":\"Palestine, State of\",\"zip_required\":false},{\"id\":170,\"name\":\"Panama\",\"zip_required\":false},{\"id\":173,\"name\":\"Papua New Guinea\",\"zip_required\":false},{\"id\":183,\"name\":\"Paraguay\",\"zip_required\":false},{\"id\":171,\"name\":\"Peru\",\"zip_required\":false},{\"id\":174,\"name\":\"Philippines\",\"zip_required\":false},{\"id\":178,\"name\":\"Pitcairn\",\"zip_required\":false},{\"id\":176,\"name\":\"Poland\",\"zip_required\":false},{\"id\":181,\"name\":\"Portugal\",\"zip_required\":false},{\"id\":179,\"name\":\"Puerto Rico\",\"zip_required\":false},{\"id\":184,\"name\":\"Qatar\",\"zip_required\":false},{\"id\":185,\"name\":\"Réunion\",\"zip_required\":false},{\"id\":186,\"name\":\"Romania\",\"zip_required\":false},{\"id\":188,\"name\":\"Russian Federation\",\"zip_required\":false},{\"id\":189,\"name\":\"Rwanda\",\"zip_required\":false},{\"id\":26,\"name\":\"Saint Barthélemy\",\"zip_required\":false},{\"id\":195,\"name\":\"Saint Helena\",\"zip_required\":false},{\"id\":118,\"name\":\"Saint Kitts And Nevis\",\"zip_required\":false},{\"id\":125,\"name\":\"Saint Lucia\",\"zip_required\":false},{\"id\":138,\"name\":\"Saint Martin\",\"zip_required\":false},{\"id\":177,\"name\":\"Saint Pierre And Miquelon\",\"zip_required\":false},{\"id\":231,\"name\":\"Saint Vincent And The Grenadines\",\"zip_required\":false},{\"id\":238,\"name\":\"Samoa\",\"zip_required\":false},{\"id\":200,\"name\":\"San Marino\",\"zip_required\":false},{\"id\":204,\"name\":\"Sao Tome and Principe\",\"zip_required\":false},{\"id\":190,\"name\":\"Saudi Arabia\",\"zip_required\":false},{\"id\":201,\"name\":\"Senegal\",\"zip_required\":false},{\"id\":187,\"name\":\"Serbia\",\"zip_required\":false},{\"id\":192,\"name\":\"Seychelles\",\"zip_required\":false},{\"id\":199,\"name\":\"Sierra Leone\",\"zip_required\":false},{\"id\":194,\"name\":\"Singapore\",\"zip_required\":false},{\"id\":206,\"name\":\"Sint Maarten\",\"zip_required\":false},{\"id\":198,\"name\":\"Slovakia\",\"zip_required\":false},{\"id\":196,\"name\":\"Slovenia\",\"zip_required\":false},{\"id\":191,\"name\":\"Solomon Islands\",\"zip_required\":false},{\"id\":202,\"name\":\"Somalia\",\"zip_required\":false},{\"id\":241,\"name\":\"South Africa\",\"zip_required\":false},{\"id\":89,\"name\":\"South Georgia and the South Sandwich Islands\",\"zip_required\":false},{\"id\":282,\"name\":\"South Sudan\",\"zip_required\":false},{\"id\":67,\"name\":\"Spain\",\"zip_required\":false},{\"id\":127,\"name\":\"Sri Lanka\",\"zip_required\":false},{\"id\":203,\"name\":\"Suriname\",\"zip_required\":false},{\"id\":197,\"name\":\"Svalbard And Jan Mayen\",\"zip_required\":false},{\"id\":207,\"name\":\"Swaziland\",\"zip_required\":false},{\"id\":193,\"name\":\"Sweden\",\"zip_required\":false},{\"id\":43,\"name\":\"Switzerland\",\"zip_required\":false},{\"id\":222,\"name\":\"Taiwan\",\"zip_required\":false},{\"id\":213,\"name\":\"Tajikistan\",\"zip_required\":false},{\"id\":223,\"name\":\"Tanzania, United Republic of\",\"zip_required\":false},{\"id\":212,\"name\":\"Thailand\",\"zip_required\":false},{\"id\":215,\"name\":\"Timor-Leste\",\"zip_required\":false},{\"id\":211,\"name\":\"Togo\",\"zip_required\":false},{\"id\":214,\"name\":\"Tokelau\",\"zip_required\":false},{\"id\":218,\"name\":\"Tonga\",\"zip_required\":false},{\"id\":220,\"name\":\"Trinidad and Tobago\",\"zip_required\":false},{\"id\":217,\"name\":\"Tunisia\",\"zip_required\":false},{\"id\":219,\"name\":\"Turkey\",\"zip_required\":false},{\"id\":216,\"name\":\"Turkmenistan\",\"zip_required\":false},{\"id\":208,\"name\":\"Turks and Caicos Islands\",\"zip_required\":false},{\"id\":221,\"name\":\"Tuvalu\",\"zip_required\":false},{\"id\":225,\"name\":\"Uganda\",\"zip_required\":false},{\"id\":224,\"name\":\"Ukraine\",\"zip_required\":false},{\"id\":2,\"name\":\"United Arab Emirates\",\"zip_required\":false},{\"id\":76,\"name\":\"United Kingdom\",\"zip_required\":false},{\"id\":226,\"name\":\"United States Minor Outlying Islands\",\"zip_required\":false},{\"id\":228,\"name\":\"Uruguay\",\"zip_required\":false},{\"id\":229,\"name\":\"Uzbekistan\",\"zip_required\":false},{\"id\":236,\"name\":\"Vanuatu\",\"zip_required\":false},{\"id\":232,\"name\":\"Venezuela, Bolivarian Republic of\",\"zip_required\":false},{\"id\":235,\"name\":\"Vietnam\",\"zip_required\":false},{\"id\":233,\"name\":\"Virgin Islands, British\",\"zip_required\":false},{\"id\":234,\"name\":\"Virgin Islands, U.S.\",\"zip_required\":false},{\"id\":237,\"name\":\"Wallis and Futuna\",\"zip_required\":false},{\"id\":65,\"name\":\"Western Sahara\",\"zip_required\":false},{\"id\":239,\"name\":\"Yemen\",\"zip_required\":false},{\"id\":242,\"name\":\"Zambia\",\"zip_required\":false},{\"id\":243,\"name\":\"Zimbabwe\",\"zip_required\":false}]"),
    allowedStatesOrProvinces: {"United Arab Emirates":["Abu Dhabi","Ajman","Dubai","Dubai","Fujairah","Ras al Khaimah","Sharjah","Umm Al Quwain"],"Afghanistan":["Badakhshan","Kabul"],"Albania":["Fier County","Tirana County"],"Armenia":["Kotayk Province","Yerevan"],"Angola":["Luanda Province"],"Argentina":["Buenos Aires","Buenos Aires","Buenos Aires Province","Cordoba","Corrientes Province","Santa Cruz Province","Santa Fe Province"],"Austria":["Carinthia","Lower Austria","Styria","Tyrol","Upper Austria","Vienna","Vorarlberg"],"Australia":["Australian Capital Territory","New South Wales","Northern Territory","Queensland","South Australia","Tasmania","Victoria","Western Australia"],"Bosnia and Herzegovina":["Federation of Bosnia and Herzegovina","Republika Srpska"],"Barbados":["Saint Michael"],"Bangladesh":["Chittagong Division","Dhaka Division","Khulna Division","Rajshahi Division","Sylhet Division"],"Belgium":["Brussels","Flanders","Wallonia"],"Bulgaria":["Plovdiv Province","Sofia City Province"],"Bahrain":["Capital Governorate","Northern Governorate","Southern Governorate"],"Benin":["Atlantique Department","Littoral Department"],"Bermuda":["Pembroke Parish"],"Bolivia":["Santa Cruz Department"],"Brazil":["Espírito Santo","Federal District","Pará","State of Bahia","State of Ceará","State of Goiás","State of Minas Gerais","State of Paraná","State of Pernambuco","State of Rio de Janeiro","State of Rio Grande do Norte","State of Rio Grande do Sul","State of Santa Catarina","State of São Paulo"],"Bahamas":["New Providence"],"Bhutan":["Thimphu"],"Botswana":["South-East District"],"Belarus":["Minsk Region","Mogilev Region"],"Canada":["Alberta","British Columbia","Manitoba","New Brunswick","Newfoundland and Labrador","Northwest Territory","Nova Scotia","Nunavut Territory","Ontario","Prince Edward Island","Quebec","Saskatchewan","Yukon"],"Congo, The Democratic Republic Of The":["Kinshasa","North-Kivu","South-Kivu"],"Switzerland":["Basel City","Canton of Bern","Geneva","Schwyz","Solothurn","Thurgau","Ticino","Vaud","Zurich"],"Côte D'Ivoire":["Abidjan Autonomous District","Comoé District"],"Chile":["Antofagasta","Bio Bio","Maule","Santiago Metropolitan Region","Valparaíso"],"Cameroon":["Centre Region","Northwest Region","Région du Nord"],"China":["Anhui","Beijing","Fujian","Gansu","Guangdong Province","Henan","Jiangsu","Shandong","Shanghai","Tianjin","Zhejiang"],"Colombia":["Antioquia","Atlantico","Bogota","Bolivar","Cundinamarca","Huila","Risaralda"],"Costa Rica":["Heredia Province","Provincia de Alajuela","San José Province"],"Curaçao":["Curaçao"],"Cyprus":["Famagusta","Larnaca","Limassol","Nicosia"],"Czech Republic":["Olomouc Region","Prague","South Moravian Region"],"Germany":["Baden-Württemberg","Bavaria","Berlin","Brandenburg","Hamburg","Hessen","Lower Saxony","Mecklenburg-Vorpommern","North Rhine-Westphalia","Rhineland-Palatinate","Saarland","Saxony","Schleswig-Holstein"],"Dominican Republic":["Distrito Nacional","Duarte Province","La Altagracia","La Romana Province"],"Algeria":["Algiers Province","Constantine Province","Jijel Province","Mila Province","Mostaganem Province","Relizane Province","Sétif Province","Skikda Province","Tlemcen Province"],"Ecuador":["Azuay","Chimborazo","Guayas","Pichincha","Santo Domingo de los Tsachilas"],"Estonia":["Harju County"],"Egypt":["Alexandria Governorate","Al-Qalyubia Governorate","Ash Sharqia Governorate","Cairo Governorate","Dakahlia Governorate","Giza Governorate","Menofia Governorate"],"Spain":["Andalusia","Aragon","Asturias","Balearic Islands","Basque Country","Canary Islands","Cantabria","Castile and León","Catalonia","Community of Madrid","Comunitat Valenciana","Galicia","Navarre"],"Ethiopia":["Addis Ababa"],"Finland":["Central Finland","North Ostrobothnia","Pirkanmaa","South Savo","Southwest Finland","Uusimaa"],"France":["Auvergne-Rhône-Alpes","Bourgogne-Franche-Comté","Brittany","Centre-Val de Loire","Grand Est","Hauts-de-France","Île-de-France","Nouvelle-Aquitaine","Occitanie","Pays de la Loire","Provence-Alpes-Côte d'Azur"],"Gabon":["Estuaire"],"United Kingdom":["England","Northern Ireland","Scotland","Wales"],"Georgia":["Adjara","Tbilisi"],"Ghana":["Central Region","Greater Accra Region","Northern Region","Western North Region"],"Guinea":["Conakry"],"Guatemala":["Guatemala Department","Quetzaltenango Department"],"Hong Kong":["Hong Kong Island"],"Honduras":["Cortés Department","Francisco Morazán Department"],"Croatia":["Osijek-Baranja County","Zagreb County"],"Indonesia":["Bali","Banten","Bengkulu","East Java","Jakarta","North Sumatra","Riau Islands","West Java","West Kalimantan","West Sumatra"],"Ireland":["County Carlow","County Cork","County Dublin","County Galway","County Kildare","County Louth","County Waterford","County Westmeath","County Wexford","County Wicklow"],"Israel":["Center District","South District","Tel Aviv District"],"India":["Andhra Pradesh","Assam","Bihar","Chandigarh","Chhattisgarh","Delhi","Goa","Gujarat","Haryana","Himachal Pradesh","Jharkhand","Karnataka","Kerala","Madhya Pradesh","Maharashtra","Meghalaya","Odisha","Puducherry","Punjab","Rajasthan","Tamil Nadu","Telangana","Telangana","Uttarakhand","Uttar Pradesh","West Bengal"],"Iraq":["Baghdad Governorate","Basra Governorate","Duhok Governorate","Karbala Governorate","Kurdistan Region"],"Italy":["Abruzzo","Apulia","Calabria","Campania","Emilia-Romagna","Friuli-Venezia Giulia","Lazio","Lombardia","Lombardy","Marche","Piedmont","Puglia","Sicilia","Sicily","Trentino-South Tyrol","Tuscany","Umbria","Veneto"],"Jamaica":["St. Andrew Parish","Westmoreland Parish"],"Jordan":["Amman Governorate"],"Japan":["Aichi","Aomori","Chiba","Fukuoka","Gifu","Gunma","Hiroshima","Ibaraki","Kanagawa","Kumamoto","Nara","Oita","Okinawa","Osaka","Saitama","Tochigi","Tokyo","Yamagata"],"Kenya":["Kiambu County","Machakos County","Mombasa County","Nairobi County","Turkana County"],"Cambodia":["Kandal Province","Phnom Penh","Siem Reap Province"],"Korea, Republic of":["Gyeonggi-do","Incheon","Seoul"],"Kuwait":["Al Asimah Governate","Al Farwaniyah Governorate","Hawalli Governorate"],"Cayman Islands":["George Town"],"Kazakhstan":["Almaty Region","Atyrau Region","Turkistan Region"],"Lebanon":["Beirut Governorate","South Governorate"],"Saint Lucia":["Gros Islet"],"Liechtenstein":["Gamprin"],"Sri Lanka":["Central Province","Eastern Province","North Central Province","North Western Province","Sabaragamuwa Province","Western Province"],"Liberia":["Montserrado"],"Lithuania":["Alytus County","Kaunas County","Vilnius County"],"Luxembourg":["Luxembourg"],"Libya":["Misrata District","Tripoli District"],"Morocco":["Casablanca-Settat","Fez-Meknès","Rabat-Salé-Kénitra","Souss-Massa","Tangier-Tétouan-Al Hoceima"],"Moldova, Republic of":["Chisinau"],"Madagascar":["Analamanga"],"North Macedonia":["Greater Skopje"],"Mali":["Bamako Capital District"],"Myanmar":["Magway Region","Tanintharyi Region","Yangon Region"],"Mauritania":["Dakhlet Nouadhibou"],"Mauritius":["Port Louis District","Rivière du Rempart District"],"Maldives":["Laamu Atoll","Malé"],"Malawi":["Northern Region","Southern Region"],"Mexico":["Aguascalientes","Baja California","Baja California Sur","Chiapas","Chihuahua","Coahuila","Colima","Durango","Guanajuato","Jalisco","Mexico City","Michoacán","Morelos","Nayarit","Nuevo Leon","Oaxaca","Puebla","Querétaro","Quintana Roo","San Luis Potosi","Sonora","State of Mexico","Tabasco","Tamaulipas","Veracruz","Yucatán"],"Malaysia":["Federal Territory of Kuala Lumpur","Johor","Kedah","Negeri Sembilan","Penang","Perak","Sarawak","Selangor"],"Mozambique":["Maputo"],"Nigeria":["Abia","Akwa Ibom","Anambra","Delta","Edo","Enugu","Federal Capital Territory","Gombe","Imo","Kaduna","Kwara","Lagos","Niger","Ogun State","Oyo","Rivers","Yobe"],"Netherlands":["Flevoland","Gelderland","Limburg","North Brabant","North Holland","Overijssel","South Holland","Utrecht"],"Norway":["Innlandet","Oslo","Rogaland","Trøndelag","Viken"],"Nepal":["Bagmati Province"],"New Zealand":["Auckland","Bay of Plenty","Canterbury","Hawke's Bay","Manawatu-Wanganui","Southland","Waikato","Wellington"],"Oman":["Ad Dakhiliyah ‍Governorate","Ad Dhahirah Governorate","Al Batinah North Governorate","Muscat Governorate"],"Panama":["Panamá Province"],"Peru":["Callao Region","Lima Province"],"Philippines":["Bicol","Cagayan Valley","Calabarzon","Central Luzon","Central Visayas","Davao Region","Ilocos Region","Metro Manila","MIMAROPA","Northern Mindanao","Western Visayas","Zamboanga Peninsula"],"Pakistan":["Balochistan","Islamabad Capital Territory","Khyber Pakhtunkhwa","Punjab","Sindh"],"Poland":["Greater Poland Voivodeship","Lesser Poland Voivodeship","Łódź Voivodeship","Lower Silesian Voivodeship","Lublin Voivodeship","Masovian Voivodeship","Podkarpackie Voivodeship","Pomeranian Voivodeship","Silesian Voivodeship"],"Puerto Rico":["Carolina","Ceiba","Río Grande","San Juan","San Lorenzo"],"Portugal":["Aveiro District","Braga","Faro District","Guarda District","Leiria District","Lisboa","Lisbon","Porto","Porto District","Setubal"],"Palau":["Koror"],"Paraguay":["Alto Paraná Department","Canindeyú"],"Qatar":["Al Daayen Municipality","Al Rayyan Municipality","Al Wakrah Municipality","Doha","Doha Municipality","Umm Salal Municipality"],"Romania":["Bucharest","Cluj County","Constanța","Timiș"],"Serbia":["Vojvodina"],"Russian Federation":["Moscow"],"Rwanda":["Kigali City","Northern Province"],"Saudi Arabia":["Al Jowf Province","Al Madinah Province","Al Qassim Province","Aseer Province","Eastern Province","Jazan Province","Makkah Province","Northern Borders Province","Riyadh Province","Tabuk Province"],"Seychelles":["Grand Anse Praslin"],"Sweden":["Dalarna County","Gavleborg County","Jonkoping County","Örebro County","Östergötland County","Skåne County","Stockholm County","Uppsala County","Västmanland County","Västra Götaland County"],"Slovenia":["Ljubljana"],"Slovakia":["Košice Region","Žilina Region"],"Sierra Leone":["Western Area"],"Senegal":["Dakar Region"],"Somalia":["Lower Juba"],"El Salvador":["Cuscatlán Department"],"Thailand":["Bangkok","Chiang Mai","Chiang Rai","Chon Buri","Nakhon Si Thammarat","Nong Khai","Nonthaburi","Pathum Thani","Udon Thani"],"Timor-Leste":["Dili"],"Tunisia":["Nabeul Governorate","Tunis Governorate"],"Turkey":["Ankara","Antalya","Hatay","İstanbul","İzmir","Kocaeli","Muğla","Sakarya","Uşak"],"Trinidad and Tobago":["Arima Borough Corporation"],"Tuvalu":["Funafuti"],"Taiwan":["Kaohsiung City","New Taipei City","Tainan City","Taipei City","Taoyuan City"],"Tanzania, United Republic of":["Arusha Region","Dar es Salam","Mjini Magharibi Region"],"Ukraine":["Kharkiv Oblast","Kyiv city","Lviv Oblast","Zhytomyr Oblast"],"Uganda":["Central Region","Eastern Region"],"United States":["Alabama","Alaska","American Samoa","Arizona","Arkansas","California","California","Colorado","Connecticut","Delaware","District of Columbia","Florida","Georgia","Guam","Hawaii","Idaho","Illinois","Indiana","Iowa","Kansas","Kentucky","Louisiana","Maine","Marshall Islands","Maryland","Massachusetts","Michigan","Minnesota","Mississippi","Missouri","Montana","Nebraska","Nevada","New Hampshire","New Jersey","New Mexico","New York","North Carolina","North Dakota","Ohio","Oklahoma","Oregon","Palau","Pennsylvania","Puerto Rico","Rhode Island","South Carolina","South Dakota","Tennessee","Texas","Texas","U.S. Virgin Islands","Utah","Vermont","Virginia","Washington","West Virginia","Wisconsin","Wyoming"],"Uruguay":["Colonia Department","Departamento de Canelones","Montevideo Department"],"Uzbekistan":["Tashkent Region"],"Venezuela, Bolivarian Republic of":["Capital District","Lara"],"Vietnam":["Binh Duong","Da Nang","Hai Phong","Hanoi","Ho Chi Minh City","Quảng Ninh","Soc Trang","Thai Nguyen"],"Yemen":["Ibb Governorate","Sana'a City"],"South Africa":["Eastern Cape","Free State","Gauteng","KwaZulu-Natal","Limpopo","Mpumalanga","North West","Western Cape"],"Zambia":["Copperbelt Province","Lusaka Province"],"Zimbabwe":["Harare Province","Mashonaland East Province","Masvingo Province","Midlands Province"]},
    country: ""
  };

  vals.available_locales = {
   allowedLanguages : [{"code":"en","name":"English"},{"code":"fr","name":"French"},{"code":"fr-CA","name":"French (Canada)"},{"code":"de","name":"German"},{"code":"ja","name":"Japanese"},{"code":"ko","name":"Korean"},{"code":"pt","name":"Portuguese"},{"code":"pt-BR","name":"Portuguese (Brazil)"},{"code":"zh-CN","name":"Simplified Chinese"},{"code":"es-ES","name":"Spanish (Spain)"},{"code":"es-US","name":"Spanish (US)"}]
 };

  vals.cdnHost = "cdn.credly.com" || document.location.host;
  vals.env = 'production';

  vals.faethm = {
    ontology_mapping_url: "https://app.faethm.ai/ontology-mapping/"
  };

  vals.frontendDebug = false

  vals.alerts = JSON.parse("{\"flashes\":[]}").flashes;

  vals.share = {
    facebookClientId: "126857714151516",
    ziprecruiter: {
      url: "https://www.ziprecruiter.com/credly",
      appName: "ZipRecruiter"
    }
  };

  vals.sso = {
      apple: {
          clientId: "com.credly.signinwithappleservicesid"
      },
      google: { clientId: "40778611547-2akc17g7rv7l4n0gfs8m9ogcn12rvl47.apps.googleusercontent.com" },
      microsoft: {
        clientId: "",
        tenantId: "54e44946-b280-4ccf-b102-2224d7008f17"
      }
  }

  vals.features = [];
  []
    .forEach(function(f) {vals.features.push(f);});
  vals.abTests = {};

  vals.config = {
    sizedImages: true,
    startingOauth: false,
    oauthAppId: ''
  };

  vals.tracking = JSON.parse(
    "{\"url\":\"https://stats.credly.com/stats/interaction\",\"refresh_url\":\"/stat_signatures/interaction\",\"request_data\":{\"session_id\":\"023f4576-0f1e-4c5e-a47b-3fe75a29d783\",\"flexible_params_auth\":[\"snapshot_json\",\"stat_object_id\",\"stat_object_type\",\"stat_type\"],\"auth_version\":\"1.0\",\"auth_key\":\"user:8d768abb-4e6a-4f38-8cb3-86dc95196f75\",\"auth_timestamp\":\"1726236354\",\"auth_signature\":\"45fed5237f097e9c1f19932662923978ecd6dba3d3e014a50eb0e91353fd283a\"}}"
  );


    vals.currentUser = { anonymous: true };

  const prerenderData = window.prerenderData = window.prerenderData || {};
    prerenderData[{"a":"/users/mikael-johansson-2/badges"}.a] =
      JSON.parse({"a":"{\"data\":[{\"id\":\"20f4aaea-770e-4e32-8cd0-f2720fb11d85\",\"expires_at_date\":\"2026-03-14\",\"issued_at_date\":\"2023-03-14\",\"issued_to\":\"Mikael Rikard Johansson Länsberg\",\"locale\":\"en\",\"public\":true,\"state\":\"accepted\",\"translate_metadata\":false,\"accepted_at\":\"2023-03-17T03:28:06.503-05:00\",\"expires_at\":\"2026-03-13T19:00:00.000-05:00\",\"issued_at\":\"2023-03-13T19:00:00.000-05:00\",\"last_updated_at\":\"2023-03-17T03:28:06.503-05:00\",\"updated_at\":\"2023-03-17T03:28:06.503-05:00\",\"earner_path\":\"/users/mikael-johansson-2\",\"earner_photo_url\":null,\"is_private_badge\":false,\"user_is_earner\":false,\"issuer\":{\"summary\":\"issued by The Linux Foundation\",\"entities\":[{\"label\":\"Issued by\",\"primary\":true,\"entity\":{\"type\":\"Organization\",\"id\":\"f4b8d042-0072-4a1a-8d00-260b513026e8\",\"name\":\"The Linux Foundation\",\"url\":\"https://www.credly.com/api/v1/organizations/f4b8d042-0072-4a1a-8d00-260b513026e8\",\"vanity_url\":\"https://www.credly.com/org/the-linux-foundation\",\"internationalize_badge_templates\":false,\"share_to_ziprecruiter\":true,\"twitter_url\":\"https://twitter.com/linuxfoundation\",\"verified\":false}}]},\"badge_template\":{\"id\":\"64567b66-def2-4c84-be6c-2586962fccd3\",\"description\":\"Earners of this designation demonstrated the skills, knowledge and competencies to perform the responsibilities of a Kubernetes Administrator. Earners demonstrated proficiency in Application Lifecycle Management, Installation, Configuration \\u0026 Validation, Core Concepts, Networking, Scheduling, Security, Cluster Maintenance, Logging / Monitoring, Storage, and Troubleshooting\",\"global_activity_url\":\"https://training.linuxfoundation.org/certification/certified-kubernetes-administrator-cka/\",\"earn_this_badge_url\":null,\"enable_earn_this_badge\":false,\"enable_detail_attribute_visibility\":true,\"name\":\"CKA: Certified Kubernetes Administrator\",\"public\":true,\"recipient_type\":\"User\",\"vanity_slug\":\"cka-certified-kubernetes-administrator\",\"show_badge_lmi\":false,\"show_skill_tag_links\":true,\"settings_enable_related_badges\":true,\"translatable\":false,\"level\":\"Intermediate\",\"time_to_earn\":\"Months\",\"cost\":\"Paid\",\"type_category\":\"Certification\",\"image\":{\"id\":\"8b8ed108-e77d-4396-ac59-2504583b9d54\",\"url\":\"https://images.credly.com/images/8b8ed108-e77d-4396-ac59-2504583b9d54/cka_from_cncfsite__281_29.png\"},\"image_url\":\"https://images.credly.com/images/8b8ed108-e77d-4396-ac59-2504583b9d54/cka_from_cncfsite__281_29.png\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/cka-certified-kubernetes-administrator\",\"owner_vanity_slug\":\"the-linux-foundation\",\"badge_template_earnable\":false,\"recommendable\":true,\"issuer\":{\"summary\":\"issued by The Linux Foundation\",\"entities\":[{\"label\":\"Issued by\",\"primary\":true,\"entity\":{\"type\":\"Organization\",\"id\":\"f4b8d042-0072-4a1a-8d00-260b513026e8\",\"name\":\"The Linux Foundation\",\"url\":\"https://www.credly.com/api/v1/organizations/f4b8d042-0072-4a1a-8d00-260b513026e8\",\"vanity_url\":\"https://www.credly.com/org/the-linux-foundation\",\"internationalize_badge_templates\":false,\"share_to_ziprecruiter\":true,\"twitter_url\":\"https://twitter.com/linuxfoundation\",\"verified\":false}}]},\"related_badge_templates\":[{\"id\":\"efc98036-fdf4-4c5c-b6ca-34e58c8d61bd\",\"name\":\"CKS: Certified Kubernetes Security Specialist\",\"image\":{\"id\":\"9945dfcb-1cca-4529-85e6-db1be3782210\",\"url\":\"https://images.credly.com/images/9945dfcb-1cca-4529-85e6-db1be3782210/kubernetes-security-specialist-logo2.png\"},\"image_url\":\"https://images.credly.com/images/9945dfcb-1cca-4529-85e6-db1be3782210/kubernetes-security-specialist-logo2.png\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/cks-certified-kubernetes-security-specialist\"},{\"id\":\"c00925a4-74e7-41cb-bbc9-eaafaf1b3d71\",\"name\":\"LFS258: Kubernetes Fundamentals\",\"image\":{\"id\":\"123746a7-fbbe-4fdd-9c0c-f0254e53292a\",\"url\":\"https://images.credly.com/images/123746a7-fbbe-4fdd-9c0c-f0254e53292a/blob\"},\"image_url\":\"https://images.credly.com/images/123746a7-fbbe-4fdd-9c0c-f0254e53292a/blob\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/lfs258-kubernetes-fundamentals\"},{\"id\":\"63527a38-01f7-468f-b015-dab43019f4cd\",\"name\":\"LFS458: Kubernetes Administration\",\"image\":{\"id\":\"20936872-ca23-44da-a4bc-db39db3468b6\",\"url\":\"https://images.credly.com/images/20936872-ca23-44da-a4bc-db39db3468b6/blob\"},\"image_url\":\"https://images.credly.com/images/20936872-ca23-44da-a4bc-db39db3468b6/blob\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/lfs458-kubernetes-administration\"}],\"alignments\":[],\"badge_template_activities\":[{\"id\":\"ffe8e8b8-40c1-42de-8c1f-22498faf6900\",\"activity_type\":\"Assessment\",\"required_badge_template_id\":null,\"title\":\"Score a passing grade on an online, proctored, performance based exam. The exam consists of a set of performance-based items (problems) to be solved in a command line, running Kubernetes.\",\"url\":\"https://training.linuxfoundation.org/certification/certified-kubernetes-administrator-cka/#domains\"}],\"endorsements\":[],\"skills\":[{\"id\":\"043c44b8-83b8-4d52-8b32-0a53d70e96d1\",\"name\":\"API objects\",\"vanity_slug\":\"api-objects\"},{\"id\":\"3fa750a1-613e-4b51-88c8-42e18a2a8ce7\",\"name\":\"Cloud\",\"vanity_slug\":\"cloud\"},{\"id\":\"07f157d5-342d-407b-9b78-ceee32aa278b\",\"name\":\"Custom Resource Definitions\",\"vanity_slug\":\"custom-resource-definitions\"},{\"id\":\"661d13df-3d75-433f-85bc-63ed774c7458\",\"name\":\"Helm\",\"vanity_slug\":\"helm\"},{\"id\":\"55e4a962-cbb2-4a9c-a452-f8aedd47ef4b\",\"name\":\"Ingress\",\"vanity_slug\":\"ingress\"},{\"id\":\"f2777485-8f70-497b-9fc4-2ed5c74c1c1f\",\"name\":\"Kubernetes\",\"vanity_slug\":\"kubernetes\"},{\"id\":\"f2a88a18-7561-47ca-bb98-63c0c395ac98\",\"name\":\"Logging and Troubleshooting\",\"vanity_slug\":\"logging-and-troubleshooting\"},{\"id\":\"23bae4b4-25e9-4a54-9323-4e19d6d0a7a8\",\"name\":\"Open Source Software\",\"vanity_slug\":\"open-source-software\"},{\"id\":\"58e2a42d-b4eb-4af7-974d-524e99b30c3c\",\"name\":\"Orchestration\",\"vanity_slug\":\"orchestration\"},{\"id\":\"69b1c2c7-ee38-4921-8d5e-fcc84a1c497f\",\"name\":\"Scheduling\",\"vanity_slug\":\"scheduling\"},{\"id\":\"db139eba-a6ef-4e44-a9d3-180563489f57\",\"name\":\"Security Policies\",\"vanity_slug\":\"security-policies\"},{\"id\":\"8576e1c9-70fc-402c-a79a-44bbd946b019\",\"name\":\"Services\",\"vanity_slug\":\"services\"},{\"id\":\"19b7554f-8bba-466d-a1e2-c4e7dc3c11bc\",\"name\":\"Site Reliability Engineer\",\"vanity_slug\":\"site-reliability-engineer\"},{\"id\":\"14b34c04-b218-4085-8426-facff7a61e50\",\"name\":\"System Administrator\",\"vanity_slug\":\"system-administrator\"},{\"id\":\"bbb2d806-696d-46e6-8aba-b5f27633e597\",\"name\":\"Volumes\",\"vanity_slug\":\"volumes\"}]},\"image\":{\"id\":\"8b8ed108-e77d-4396-ac59-2504583b9d54\",\"url\":\"https://images.credly.com/images/8b8ed108-e77d-4396-ac59-2504583b9d54/cka_from_cncfsite__281_29.png\"},\"image_url\":\"https://images.credly.com/images/8b8ed108-e77d-4396-ac59-2504583b9d54/cka_from_cncfsite__281_29.png\",\"evidence\":[{\"id\":\"9bf7a6e5-743c-46fe-a138-ab8f1a6ca24a\",\"type\":\"PlainTextEvidence\",\"title\":\"Certificate ID Number\",\"description\":\"LF-bv8xpen5hz\"}],\"recommendations\":[]},{\"id\":\"062ae104-f532-43d0-b3bd-b6599dd03e2c\",\"expires_at_date\":\"2026-08-30\",\"issued_at_date\":\"2024-08-29\",\"issued_to\":\"Mikael Rikard Johansson Länsberg\",\"locale\":\"en\",\"public\":true,\"state\":\"accepted\",\"translate_metadata\":false,\"accepted_at\":\"2024-08-29T03:36:04.318-05:00\",\"expires_at\":\"2026-08-29T19:00:00.000-05:00\",\"issued_at\":\"2024-08-28T19:00:00.000-05:00\",\"last_updated_at\":\"2024-08-29T03:36:04.338-05:00\",\"updated_at\":\"2024-08-29T03:36:04.475-05:00\",\"earner_path\":\"/users/mikael-johansson-2\",\"earner_photo_url\":null,\"is_private_badge\":false,\"user_is_earner\":false,\"issuer\":{\"summary\":\"issued by The Linux Foundation\",\"entities\":[{\"label\":\"Issued by\",\"primary\":true,\"entity\":{\"type\":\"Organization\",\"id\":\"f4b8d042-0072-4a1a-8d00-260b513026e8\",\"name\":\"The Linux Foundation\",\"url\":\"https://www.credly.com/api/v1/organizations/f4b8d042-0072-4a1a-8d00-260b513026e8\",\"vanity_url\":\"https://www.credly.com/org/the-linux-foundation\",\"internationalize_badge_templates\":false,\"share_to_ziprecruiter\":true,\"twitter_url\":\"https://twitter.com/linuxfoundation\",\"verified\":false}}]},\"badge_template\":{\"id\":\"c54d5824-9af2-4bf4-8ce0-323e35167701\",\"description\":\"Earners of this designation demonstrated a basic knowledge of Kubernetes and cloud-native technologies, including how to deploy an application using basic kubectl commands, the architecture of Kubernetes (containers, pods, nodes, clusters), understanding the cloud-native landscape and projects (storage, networking, GitOps, service mesh), and understanding the principles of cloud-native security.\",\"global_activity_url\":\"https://training.linuxfoundation.org/certification/kubernetes-cloud-native-associate/\",\"earn_this_badge_url\":null,\"enable_earn_this_badge\":false,\"enable_detail_attribute_visibility\":true,\"name\":\"KCNA: Kubernetes and Cloud Native Associate\",\"public\":true,\"recipient_type\":\"User\",\"vanity_slug\":\"kcna-kubernetes-and-cloud-native-associate\",\"show_badge_lmi\":false,\"show_skill_tag_links\":true,\"settings_enable_related_badges\":true,\"translatable\":false,\"level\":\"Foundational\",\"time_to_earn\":\"Weeks\",\"cost\":\"Paid\",\"type_category\":\"Certification\",\"image\":{\"id\":\"f28f1d88-428a-47f6-95b5-7da1dd6c1000\",\"url\":\"https://images.credly.com/images/f28f1d88-428a-47f6-95b5-7da1dd6c1000/KCNA_badge.png\"},\"image_url\":\"https://images.credly.com/images/f28f1d88-428a-47f6-95b5-7da1dd6c1000/KCNA_badge.png\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/kcna-kubernetes-and-cloud-native-associate\",\"owner_vanity_slug\":\"the-linux-foundation\",\"badge_template_earnable\":false,\"recommendable\":true,\"issuer\":{\"summary\":\"issued by The Linux Foundation\",\"entities\":[{\"label\":\"Issued by\",\"primary\":true,\"entity\":{\"type\":\"Organization\",\"id\":\"f4b8d042-0072-4a1a-8d00-260b513026e8\",\"name\":\"The Linux Foundation\",\"url\":\"https://www.credly.com/api/v1/organizations/f4b8d042-0072-4a1a-8d00-260b513026e8\",\"vanity_url\":\"https://www.credly.com/org/the-linux-foundation\",\"internationalize_badge_templates\":false,\"share_to_ziprecruiter\":true,\"twitter_url\":\"https://twitter.com/linuxfoundation\",\"verified\":false}}]},\"related_badge_templates\":[{\"id\":\"935f1f2a-29ee-4f56-a4d6-e40e93ecd6fd\",\"name\":\"PCA: Prometheus Certified Associate\",\"image\":{\"id\":\"c34436dc-1cfd-4125-a862-35f9c86ca17f\",\"url\":\"https://images.credly.com/images/c34436dc-1cfd-4125-a862-35f9c86ca17f/image.png\"},\"image_url\":\"https://images.credly.com/images/c34436dc-1cfd-4125-a862-35f9c86ca17f/image.png\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/pca-prometheus-certified-associate\"},{\"id\":\"bbeb2026-8eee-4f12-94fd-0cc15cf9231b\",\"name\":\"LFS250: Kubernetes and Cloud Native Essentials\",\"image\":{\"id\":\"7404ca0d-98e1-48b6-a2a3-de8d7dcd85b5\",\"url\":\"https://images.credly.com/images/7404ca0d-98e1-48b6-a2a3-de8d7dcd85b5/blob\"},\"image_url\":\"https://images.credly.com/images/7404ca0d-98e1-48b6-a2a3-de8d7dcd85b5/blob\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/lfs250-kubernetes-and-cloud-native-essentials\"},{\"id\":\"1893a11c-6a1e-4e99-a655-c39c2297d0e9\",\"name\":\"Speaker: Open Source Summit Europe 2020\",\"image\":{\"id\":\"64fafd24-2358-4434-9a6a-cd6700ae458e\",\"url\":\"https://images.credly.com/images/64fafd24-2358-4434-9a6a-cd6700ae458e/LF_Events20_DigitalBadges_FNL_OSS_EU_Speaker.png\"},\"image_url\":\"https://images.credly.com/images/64fafd24-2358-4434-9a6a-cd6700ae458e/LF_Events20_DigitalBadges_FNL_OSS_EU_Speaker.png\",\"url\":\"https://www.credly.com/org/the-linux-foundation/badge/speaker-open-source-summit-europe-2020\"}],\"alignments\":[],\"badge_template_activities\":[{\"id\":\"847ee230-8782-430f-94a4-b55ed70b236f\",\"activity_type\":\"Assessment\",\"required_badge_template_id\":null,\"title\":\"Score a passing grade on a 90 minute online, proctored, multiple-choice exam\",\"url\":\"https://training.linuxfoundation.org/certification/kubernetes-cloud-native-associate/\"}],\"endorsements\":[],\"skills\":[{\"id\":\"15cf15e4-ef21-4f1f-b650-9a7c7249ff93\",\"name\":\"Application Containers\",\"vanity_slug\":\"application-containers\"},{\"id\":\"fe755636-ff24-4721-95c8-8a16ce83208a\",\"name\":\"Cloud + Cloud Native\",\"vanity_slug\":\"cloud-cloud-native\"},{\"id\":\"4b13d832-d606-429f-ad83-df579f629706\",\"name\":\"Clusters\",\"vanity_slug\":\"clusters\"},{\"id\":\"9c433116-d986-48fa-9346-8f8c5adbcdf2\",\"name\":\"Git Ops\",\"vanity_slug\":\"git-ops\"},{\"id\":\"f2777485-8f70-497b-9fc4-2ed5c74c1c1f\",\"name\":\"Kubernetes\",\"vanity_slug\":\"kubernetes\"},{\"id\":\"a7fa8ebb-5532-48f5-b2bb-3385a4ac757e\",\"name\":\"Pods\",\"vanity_slug\":\"pods\"},{\"id\":\"bde085ad-dfbd-4fbc-9168-5124e058038b\",\"name\":\"Service Mesh\",\"vanity_slug\":\"service-mesh\"}]},\"image\":{\"id\":\"f28f1d88-428a-47f6-95b5-7da1dd6c1000\",\"url\":\"https://images.credly.com/images/f28f1d88-428a-47f6-95b5-7da1dd6c1000/KCNA_badge.png\"},\"image_url\":\"https://images.credly.com/images/f28f1d88-428a-47f6-95b5-7da1dd6c1000/KCNA_badge.png\",\"evidence\":[{\"id\":\"95779abc-c4a0-4729-858e-fc78b3499ab1\",\"type\":\"PlainTextEvidence\",\"title\":\"Certificate ID Number\",\"description\":\"LF-f77jxibne9\"}],\"recommendations\":[]}],\"metadata\":{\"count\":2,\"current_page\":1,\"total_count\":2,\"total_pages\":1,\"per\":48,\"previous_page_url\":null,\"next_page_url\":null,\"key\":{\"page\":1,\"sort\":\"most_popular\",\"user_id\":\"mikael-johansson-2\"}}}"}.a);
      prerenderData[{"a":"/users/mikael-johansson-2/badges"}.a].list =
        {"a":"general"}.a;
    prerenderData[{"a":"/users/mikael-johansson-2"}.a] =
      JSON.parse({"a":"{\"data\":{\"id\":\"mikael-johansson-2\",\"first_name\":\"Mikael\",\"middle_name\":null,\"last_name\":\"Johansson\",\"vanity_slug\":\"mikael-johansson-2\",\"photo\":null,\"photo_url\":null,\"synthetic_id\":\"a52c3dc9-d0b3-4f23-b319-14756ed00dd0\",\"bio\":null,\"current_position_name\":\"Site Reliability Engineer\",\"current_organization_name\":\"Etraveli Group\",\"messaging_enabled\":false,\"website_url\":null,\"city\":\"Gothenburg\",\"state\":null,\"profile_is_public\":true,\"vanity_url\":\"https://www.credly.com/users/mikael-johansson-2\",\"social_profiles\":[{\"profile_url\":null,\"name\":\"Facebook\"},{\"profile_url\":null,\"name\":\"LinkedIn\"},{\"profile_url\":null,\"name\":\"Twitter\"}],\"skills\":[],\"experiences\":[]},\"metadata\":{}}"}.a);
})();


//]]>
</script><script src="https://cdn.credly.com/assets/application.legacy-browser-864855c70504f79c582e4d2832337221427fbcf00a737adbd2547bbcc05a41bd.js" crossorigin="anonymous"></script>
</body>
</html>
`

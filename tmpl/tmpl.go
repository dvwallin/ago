package tmpl

// Header contains html code and placeholder tags
const Header = `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title>[[TITLE]]</title><meta name="viewport" content="width=device-width, initial-scale=1"><meta name="description" content="[[DESCRIPTION]]"><meta name="keywords" content="[[TAGS]]"><style>%%STYLE%%</style><link rel="alternate" type="application/rss+xml" title="RSS Feed for [[DOMAIN]]" href="[[PROTOCOL]]://[[DOMAIN]]/ago.rss" /><link rel="alternate" type="application/atom+xml" title="Atom Feed for [[DOMAIN]]" href="[[PROTOCOL]]://[[DOMAIN]]/ago.atom" /></head><body><header><h2><a href="[[PROTOCOL]]://[[DOMAIN]]">[[TITLE]]</a></h2><p><em>[[INTRO]]</em></p><hr /><nav><ul><li><a href="[[PROTOCOL]]://[[DOMAIN]]/all_entries.html">View all entries</a></li><li><a href="[[PROTOCOL]]://[[DOMAIN]]/ago.atom">Atom feed</a></li><li><a href="[[PROTOCOL]]://[[DOMAIN]]/ago.rss">RSS feed</a></li></ul></nav><hr /></header>`

// Footer contains the ending of the html
const Footer = `<footer>generated with the <a href="https://ago.ofnir.xyz">ago blog</a> script. source code located at <a href="https://github.com/dvwallin/ago">GitHub</a>.</footer></body></html>`

// Style is the stylesheet to be used
const Style = `body{font-family:system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;line-height:1.4;max-width:800px;margin:20px auto;padding:0 10px;color:var(--text-main);background:var(--background-body);text-rendering:optimizeLegibility}button,input,textarea{transition:background-color var(--animation-duration) linear, border-color var(--animation-duration) linear, color var(--animation-duration) linear, box-shadow var(--animation-duration) linear, transform var(--animation-duration) ease}h1{font-size:2.2em;margin-top:0}h1,h2,h3,h4,h5,h6{margin-bottom:12px}h1,h2,h3,h4,h5,h6,strong{color:var(--text-bright)}b,h1,h2,h3,h4,h5,h6,strong,th{font-weight:600}q:before{content:none}q:after{content:none}blockquote,q{border-left:4px solid var(--focus);margin:1.5em 0;padding:0.5em 1em;font-style:italic}blockquote > footer{font-style:normal;border:0}blockquote cite{font-style:normal}address{font-style:normal}a[href^='mailto']::before{content:'📧 '}a[href^='tel']::before{content:'📞 '}a[href^='sms']::before{content:'💬 '}mark{background-color:var(--highlight);border-radius:2px;padding:0 2px;color:#000000}button,input[type='submit'],input[type='button'],input[type='checkbox'],input[type='range'],input[type='radio'],select{cursor:pointer}input:not([type='checkbox']):not([type='radio']),select{display:block}button,input,select,textarea{color:var(--form-text);background-color:var(--background);font-family:inherit;font-size:inherit;margin-right:6px;margin-bottom:6px;padding:10px;border:none;border-radius:6px;outline:none}button,input,select,textarea{-webkit-appearance:none}textarea{margin-right:0;width:100%;box-sizing:border-box;resize:vertical}select{background:var(--background) var(--select-arrow) calc(100% - 12px) 50% / 12px no-repeat;padding-right:35px}select::-ms-expand{display:none}select[multiple]{padding-right:10px;background-image:none;overflow-y:auto}button,input[type='submit'],input[type='button']{padding-right:30px;padding-left:30px}button:hover,input[type='submit']:hover,input[type='button']:hover{background:var(--button-hover)}button:focus,input:focus,select:focus,textarea:focus{box-shadow:0 0 0 2px var(--focus)}input[type='checkbox'],input[type='radio']{position:relative;width:14px;height:14px;display:inline-block;vertical-align:middle;margin:0 2px 0 0}input[type='radio']{border-radius:50%}input[type='checkbox']:checked,input[type='radio']:checked{background:var(--button-hover)}input[type='checkbox']:checked::before,input[type='radio']:checked::before{content:'•';display:block;position:absolute;left:50%;top:50%;transform:translateX(-50%) translateY(-50%)}input[type='checkbox']:checked::before{content:'✔';transform:translateY(-50%) translateY(0.5px) translateX(-6px)}input[type='checkbox']:active,input[type='radio']:active,input[type='submit']:active,input[type='button']:active,input[type='range']:active,button:active{transform:translateY(2px)}button:disabled,input:disabled,select:disabled,textarea:disabled{cursor:not-allowed;opacity:0.5}::-webkit-input-placeholder{color:var(--form-placeholder)}::-moz-placeholder{color:var(--form-placeholder)}::-ms-input-placeholder{color:var(--form-placeholder)}::placeholder{color:var(--form-placeholder)}fieldset{border:1px var(--focus) solid;border-radius:6px;margin:0 0 6px;padding:10px}legend{font-size:0.9em;font-weight:600}input[type='range']{margin:10px 0;padding:10px 0;background:transparent}input[type='range']:focus{outline:none}input[type='range']::-webkit-slider-runnable-track{width:100%;height:9.5px;transition:0.2s;background:var(--background);border-radius:3px}input[type='range']::-webkit-slider-thumb{box-shadow:0 1px 1px #000000, 0 0 1px #0d0d0d;height:20px;width:20px;border-radius:50%;background:var(--border);-webkit-appearance:none;margin-top:-7px}input[type='range']:focus::-webkit-slider-runnable-track{background:var(--background)}input[type='range']::-moz-range-track{width:100%;height:9.5px;transition:0.2s;background:var(--background);border-radius:3px}input[type='range']::-moz-range-thumb{box-shadow:1px 1px 1px #000000, 0 0 1px #0d0d0d;height:20px;width:20px;border-radius:50%;background:var(--border)}input[type='range']::-ms-track{width:100%;height:9.5px;background:transparent;border-color:transparent;border-width:16px 0;color:transparent}input[type='range']::-ms-fill-lower{background:var(--background);border:0.2px solid #010101;border-radius:3px;box-shadow:1px 1px 1px #000000, 0 0 1px #0d0d0d}input[type='range']::-ms-fill-upper{background:var(--background);border:0.2px solid #010101;border-radius:3px;box-shadow:1px 1px 1px #000000, 0 0 1px #0d0d0d}input[type='range']::-ms-thumb{box-shadow:1px 1px 1px #000000, 0 0 1px #0d0d0d;border:1px solid #000000;height:20px;width:20px;border-radius:50%;background:var(--border)}input[type='range']:focus::-ms-fill-lower{background:var(--background)}input[type='range']:focus::-ms-fill-upper{background:var(--background)}a{text-decoration:none;color:var(--links)}a{color:#4360a2;}a:hover{text-decoration:underline}code,samp,time{background:var(--background);color:var(--code);padding:2.5px 5px;border-radius:6px;font-size:1em}pre > code{padding:10px;display:block;overflow-x:auto}var{color:var(--variable);font-style:normal;font-family:monospace}kbd{background:var(--background);border:1px solid var(--border);border-radius:2px;color:var(--text-main);padding:2px 4px}img{max-width:100%;height:auto}hr{border:none;border-top:1px solid var(--border)}table{border-collapse:collapse;margin-bottom:10px;width:100%}td,th{padding:6px;text-align:left}thead{border-bottom:1px solid var(--border)}tfoot{border-top:1px solid var(--border)}tbody tr:nth-child(even){background-color:var(--background-alt)}::-webkit-scrollbar{height:10px;width:10px}::-webkit-scrollbar-track{background:var(--background);border-radius:6px}::-webkit-scrollbar-thumb{background:var(--scrollbar-thumb);border-radius:6px}::-webkit-scrollbar-thumb:hover{background:var(--scrollbar-thumb-hover)}::-moz-selection{background-color:var(--selection)}::selection{background-color:var(--selection)}details{display:flex;flex-direction:column;align-items:flex-start;background-color:var(--background-alt);padding:10px 10px 0;margin:1em 0;border-radius:6px;overflow:hidden}details[open]{padding:10px}details > :last-child{margin-bottom:0}details[open] summary{margin-bottom:10px}summary{display:list-item;background-color:var(--background);padding:10px;margin:-10px -10px 0}details > :not(summary){margin-top:0}summary::-webkit-details-marker{color:var(--text-main)}footer{border-top:1px solid var(--background);padding-top:10px;font-size:0.8em;color:var(--text-muted)}`

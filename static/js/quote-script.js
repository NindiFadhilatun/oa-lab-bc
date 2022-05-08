function splitWords() {
  let quote = document.querySelector("blockquote q"),quotewords,wordCount;
  quote.innerText.replace(/(<([^>]+)>)/ig,"");
  quotewords = quote.innerText.split(" ");
  wordCount = quotewords.length;
  quote.innerHTML = "";
  for (let i=0; i < wordCount; i++) {quote.innerHTML += "<span>"+quotewords[i]+"</span>";
    if (i < quotewords.length - 1) {quote.innerHTML += " ";}}
  quotewords = document.querySelectorAll("blockquote q span");
  fadeWords(quotewords);}
function getRandom(min,max) {return Math.random() * (max - min) + min;}
function fadeWords(quotewords) {
  Array.prototype.forEach.call(quotewords,function(word) {
    word.animate([{
          opacity: 0,
          filter: "blur("+getRandom(2,5)+"px)"},{
          opacity: 1,
          filter: "blur(0px)"}],
        {duration: 750,
          delay: getRandom(300,1750),
          fill: 'forwards'});})}
splitWords();
window.addEventListener('DOMContentLoaded',function() {
  const quote = document.getElementById("quote");
  let r = document.querySelector(':root'),
      numWords = quote.textContent.split(' ').length;
  if (numWords > 80) {r.style.setProperty('--font-size','1.15rem');
  } else if (numWords > 60) {r.style.setProperty('--font-size','1.4rem');
  } else if (numWords > 40) {r.style.setProperty('--font-size','1.65rem');}});
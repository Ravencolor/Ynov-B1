'use strict';   // Mode strict du JavaScript

/*************************************************************************************************/
/* ****************************************** DONNEES ****************************************** */
/*************************************************************************************************/

var slides =
[
	{image: "images/1.jpg", legende: "Street Art"},
	{image: "images/2.jpg", legende: "image 2"},
	{image: "images/3.jpg", legende: "Colorful Building"},
	{image: "images/4.jpg", legende: "Skyscrapers"},
	{image: "images/5.jpg", legende: "Citu by night"},
	{image: "images/6.jpg", legende: "Tour Eiffel la nuit"}
];

			  //console.log(slides);

var indiceImg;
var timer = null;

/*************************************************************************************************/
/* ***************************************** FONCTIONS ***************************************** */
/*************************************************************************************************/
function refreshSlider()

{
	var myimg = document.querySelector('img');

	var mycaption = document.querySelector('figcaption');

	myimg.src = slides[indiceImg].image;

	mycaption.textContent = slides[indiceImg].legende;

}

function bareOutils()
{
	var myicon = document.querySelector('#toolbar-toggle > i');

	myicon.classList.toggle('fa-arrow-down');
	myicon.classList.toggle('fa-arrow-right');

	document.querySelector('ul').classList.toggle('hide');
}

function onClickSliderRandom()
{

	var posImg;

	do
	{
		posImg = getRandomInteger(0,slides.length-1);

	}
	while (posImg == indiceImg);

	indiceImg = posImg;

	refreshSlider();
}

function onClickSliderPrecedent()
{
		console.log('avant',indiceImg);

		indiceImg--;

		if(indiceImg == -1)
		{
			indiceImg = slides.length - 1;
		}

		console.log('apres',indiceImg);

		refreshSlider();

}

 function onClickSliderSuivant()

{
	indiceImg ++;

	if(indiceImg == slides.length)
		{
			indiceImg = 0;
		}
	refreshSlider();

}

function onClickSliderToogle()
{

	document.querySelector('#slider-toggle i').classList.toggle('fa-pause');
	document.querySelector('#slider-toggle i').classList.toggle('fa-play');
	document.querySelector('#slider-toggle').title = "pause";

	if(timer == null)
	{

	timer = setInterval(onClickSliderSuivant,1000);
	}

	else 

	{
		document.querySelector('#slider-toggle').title = "play";
		
		clearInterval(timer);
		timer = null;
	}

}

function onSlideKeyup(e)
{
	switch (e.keyCode)
	{
		case 32:
		onClickSliderToogle();
		break;
		case 37:
		onClickSliderPrecedent();
		break;
		case 39:
		onClickSliderSuivant();
		break;
		default: console.log('ereur clavier');
	}
}

/*************************************************************************************************/
/* ************************************** CODE PRINCIPAL *************************************** */
/*************************************************************************************************/
indiceImg = 0;

refreshSlider();

var mybareOutils = document.querySelector('#toolbar-toggle');
var btnPrecedent = document.querySelector('#slider-previous');
var btnSuivant = document.querySelector('#slider-next');

mybareOutils.addEventListener('click', bareOutils);

btnPrecedent.addEventListener('click', onClickSliderPrecedent);

btnSuivant.addEventListener('click', onClickSliderSuivant);
 
document.querySelector('#slider-random').addEventListener('click',onClickSliderRandom);

document.querySelector('#slider-toggle').addEventListener('click',onClickSliderToogle);
document.addEventListener('keyup',onSlideKeyup);

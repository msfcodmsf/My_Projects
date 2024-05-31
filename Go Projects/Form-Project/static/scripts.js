// ... other JavaScript ...
document.addEventListener('DOMContentLoaded', () => {
    const themeToggle = document.getElementById('themeToggle');
    const themeIcon = document.getElementById('themeIcon');

    themeToggle.addEventListener('click', () => {
        document.body.classList.toggle('night-mode');

        if (document.body.classList.contains('night-mode')) {
            themeIcon.classList.remove('fa-sun');
            themeIcon.classList.add('fa-moon');
        } else {
            themeIcon.classList.remove('fa-moon');
            themeIcon.classList.add('fa-sun');
        }
    });
});




// ... (gece/gündüz modu kodları)

const languages = ["Python", "JavaScript", "Java", "C#", "C++", "PHP", "Ruby", "Swift", "Go", "Kotlin", "Rust"];
const searchBox = document.getElementById('searchBox');
const resultsContainer = document.getElementById('resultsContainer');
const selectedLanguages = [];
const selectedLanguagesSet = new Set();
const selectedLanguagesContainer = document.getElementById('selectedLanguagesContainer');
const selectedLanguagesList = document.getElementById('selectedLanguagesList');
const clearButton = document.getElementById('clearButton');

function searchLanguages() {
    const searchTerm = searchBox.value.toLowerCase();
    const filteredLanguages = languages.filter(lang => 
        lang.toLowerCase().includes(searchTerm) && !selectedLanguages.includes(lang) // Seçilmemiş dilleri filtrele
    );

    resultsContainer.innerHTML = '';
    filteredLanguages.forEach(lang => {
        const listItem = document.createElement('li');
        listItem.textContent = lang;
        listItem.onclick = () => selectLanguage(lang);
        resultsContainer.appendChild(listItem);
    });
}

function selectLanguage(language) {
    if (!selectedLanguagesSet.has(language)) {
        selectedLanguagesSet.add(language);
        updateSelectedLanguages();
    }
}

function updateSelectedLanguages() {
    resultsContainer.innerHTML = ''; // Önceki sonuçları temizle

    // Seçilen dilleri bir kez yazdır
    const selectedLanguagesHeader = document.createElement('div');
    selectedLanguagesHeader.textContent = "Seçilen Diller:";
    resultsContainer.appendChild(selectedLanguagesHeader);

    // Seçilen dillerin listesini oluştur
    const selectedLanguagesList = document.createElement('ul');
    selectedLanguagesSet.forEach(lang => {
        const listItem = document.createElement('li');
        listItem.textContent = lang;
        selectedLanguagesList.appendChild(listItem);
    });
    resultsContainer.appendChild(selectedLanguagesList);
}
updateSelectedLanguages();

clearButton.addEventListener('click', () => {
    selectedLanguagesSet.clear(); // Set'i temizle
    updateSelectedLanguages();
    searchBox.value = ''; // Arama kutusunu temizle
    searchLanguages(); // Arama sonuçlarını güncelle
});
  

(function(window, document, undefined){
    "use strict";
    var nightMode = document.cookie.indexOf("nightMode=true") !== -1;
    var lightMode = document.cookie.indexOf("nightMode=false") !== -1;
    if (nightMode){
      document.body.classList.add("night-mode");
    } else {
      document.body.classList.add("light-mode");
    }
    
    const userPrefersDark = window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;
    const userPrefersLight = window.matchMedia && window.matchMedia("(prefers-color-scheme: light)").matches;

    if (!lightMode && userPrefersDark){
        document.body.classList.add("night-mode");
    }
    if (!nightMode && userPrefersLight){
        document.body.classList.add("light-mode");
    }
})(window, document);

(function(window, document, undefined){
    "use strict";
    var nav = document.querySelector(".theme-mode");
    nav.innerHTML += '<span id="night-mode"><a role="button" title="nightMode" href="javascript:void(0);">🌓</a></span>';
    var nightMode = document.querySelector("#night-mode");
    nightMode.addEventListener("click", function(event){
        event.preventDefault();
        document.body.classList.toggle("light-mode");
        document.body.classList.toggle("night-mode");
        if (document.body.classList.contains("night-mode")){
            document.cookie = "nightMode=true; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/; secure;";
        } else {
            document.cookie = "nightMode=false; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/; secure;";
        }
    }, false);
})(window, document);

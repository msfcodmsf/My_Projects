document.addEventListener('DOMContentLoaded', function () {
  // Tema değiştirme düğmesi ve ikonu, sayfa ve body elementleri
  const themeToggle = document.getElementById('themeToggle');
  const themeIcon = document.getElementById('themeIcon');
  const body = document.body;

  // Temayı ayarlayan fonksiyon
  const setTheme = (theme) => {
    if (theme === 'night') {
      body.classList.add('night-mode');
      themeIcon.classList.replace('fa-sun', 'fa-moon');
    } else if (theme === 'light') {
      body.classList.remove('night-mode');
      themeIcon.classList.replace('fa-moon', 'fa-sun');
    }
  };

  // Tema değiştirme düğmesine tıklama olayı
  themeToggle.addEventListener('click', function () {
    let currentTheme = 'day';
    if (body.classList.contains('night-mode')) {
      currentTheme = 'night';
    } else if (body.style.getPropertyValue('--background-color') === 'rgb(255, 255, 255)') {
      currentTheme = 'light';
    }
    // Tema durumuna göre temayı değiştir ve localStorage'a kaydet
    if (currentTheme === 'day') {
      setTheme('night');
      localStorage.setItem('theme', 'night');
    } else if (currentTheme === 'night') {
      setTheme('light');
      localStorage.setItem('theme', 'light');
    }
  });
  // Kayıtlı temayı yükle veya varsayılan temayı kullan
  const savedTheme = localStorage.getItem('theme') || 'day';
  setTheme(savedTheme);
});

// Dil arama fonksiyonu
function searchLanguages() {
  // Arama kutusu, filtre ve sonuçları al
  const input = document.getElementById('searchBox');
  const filter = input.value.toLowerCase();
  const nodes = document.getElementsByClassName('post');

  // Arama sonuçlarını filtrele ve g ster
  for (let i = 0; i < nodes.length; i++) {
    if (nodes[i].innerText.toLowerCase().includes(filter)) {
      nodes[i].style.display = "flex";
    } else {
      nodes[i].style.display = "none";
    }
  }
}

const selectedLanguages = [];
const selectedLanguagesSet = new Set();
const selectedLanguagesContainer = document.getElementById('selectedLanguagesContainer');

// Dil arama fonksiyonu (filtreleme ve sonuçları g sterme)
function searchLanguages() {
  const searchTerm = searchBox.value.toLowerCase();
  const filteredLanguages = languages.filter(lang =>
    lang.toLowerCase().includes(searchTerm) && !selectedLanguages.includes(lang) // Seçilmemiş dilleri filtrele
  );
  // Sonuçları g ster
  resultsContainer.innerHTML = '';
  filteredLanguages.forEach(lang => {
    const listItem = document.createElement('li');
    listItem.textContent = lang;
    listItem.onclick = () => selectLanguage(lang);
    resultsContainer.appendChild(listItem);
  });
}

// Dil seçme fonksiyonu
function selectLanguage(language) {
  if (!selectedLanguagesSet.has(language) && selectedLanguages.length < 3) {
    selectedLanguagesSet.add(language);
    updateSelectedLanguages();
  }
}

// Seçilen dilleri g ncelle
function updateSelectedLanguages() {
  selectedLanguagesList.innerHTML = ''; // Önceki seçilen dilleri temizle

  // Seçilen dillerin listesini g ncelle
  selectedLanguagesSet.forEach(lang => {
    const listItem = document.createElement('li');
    listItem.textContent = lang;

    // Kategori silme işlevselliği ekleyin
    listItem.onclick = () => {
      selectedLanguagesSet.delete(lang);
      updateSelectedLanguages();
    };

    selectedLanguagesList.appendChild(listItem);
  });
}

// Temizleme d ğmesine tıklama olayı
clearButton.addEventListener('click', () => {
  selectedLanguagesSet.clear(); // Set'i temizle
  updateSelectedLanguages();
  searchBox.value = ''; // Arama kutusunu temizle
  searchLanguages(); // Arama sonuçlarını g ncelle
});

// Dil arama kutusuna herhangi bir değişiklik yapıldığında aramayı başlatın
searchBox.addEventListener('input', searchLanguages);

// Sayfa y klendiğinde seçilen dilleri g ncelle
updateSelectedLanguages();


// Kullanıcı tercihlerine ve çerezlere dayalı olarak tema modunu belirler ve ayarlar
(function (window, document, undefined) {
  "use strict";
  // Tarayıcıda kayıtlı "nightMode" çerezinin varlığını kontrol eder
  var nightMode = document.cookie.indexOf("nightMode=true") !== -1;
  // Tarayıcıda kayıtlı "nightMode" çerezinin olmadığını ve "lightMode" çerezinin varlığını kontrol eder
  var lightMode = document.cookie.indexOf("nightMode=false") !== -1;
  // Eğer "nightMode" çerezi varsa, body elementine "night-mode" sınıfını ekler
  if (nightMode) {
    document.body.classList.add("night-mode");
  } else {
    // Eğer "nightMode" çerezi yoksa, body elementine varsayılan olarak "light-mode" sınıfını ekler
    document.body.classList.add("light-mode");
  }

  // Kullanıcının tarayıcı tercihlerine göre temayı belirler
  const userPrefersDark = window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;
  const userPrefersLight = window.matchMedia && window.matchMedia("(prefers-color-scheme: light)").matches;

  // Eğer tarayıcıda "nightMode" çerezi yok ve kullanıcı karanlık modu tercih ediyorsa, body elementine "night-mode" sınıfını ekler
  if (!lightMode && userPrefersDark) {
    document.body.classList.add("night-mode");
  }
  // Eğer tarayıcıda "nightMode" çerezi yok ve kullanıcı açık modu tercih ediyorsa, body elementine "light-mode" sınıfını ekler
  if (!nightMode && userPrefersLight) {
    document.body.classList.add("light-mode");
  }
})(window, document);

// Tema değiştirme işlevselliği ekler
(function (window, document, undefined) {
  "use strict";
  // Tema değiştirme düğmesini ve ikonunu seçer
  var nav = document.querySelector(".theme-mode");
  // Tema değiştirme düğmesine bir gece modu düğmesi ekler
  nav.innerHTML += '<span id="night-mode"><a role="button" title="nightMode" href="javascript:void(0);">🌓</a></span>';
  var nightMode = document.querySelector("#night-mode");
  // Gece modu düğmesine tıklama olayı ekler
  nightMode.addEventListener("click", function (event) {
    event.preventDefault();
    // Body elementinden gece modu sınıfını kaldırır veya ekler
    document.body.classList.toggle("light-mode");
    document.body.classList.toggle("night-mode");
    // Body'nin sınıflarına göre "nightMode" çerezini ayarlar
    if (document.body.classList.contains("night-mode")) {
      document.cookie = "nightMode=true; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/; secure;";
    } else {
      document.cookie = "nightMode=false; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/; secure;";
    }
  }, false);
})(window, document);

// Sayfa yüklendiğinde tartışma başlığı, etiketler ve içerik için CKEditor'ı başlatır ve gönderme işlevselliği ekler
document.addEventListener('DOMContentLoaded', function () {
  // HTML'de tanımlı başlık, etiketler ve gönderme düğmesini seçer
  const baslik = document.getElementById("baslik");
  const etiketler = document.getElementById("etiketler");
  const gonderButton = document.getElementById("gonder");

  // CKEditor'ı başlatır
  CKEDITOR.replace('editor');

  // Gönderme düğmesine tıklama olayı ekler
  gonderButton.addEventListener("click", () => {
    const editorData = CKEDITOR.instances.editor.getData(); // CKEditor içeriğini alır
    // Editor içeriğini sunucuya gönderir (örneğin: başlık, etiket, içerik)
    alert("Tartışma başlatıldı! İçerik:\n" + editorData);
  });
});

// Belirli sınıfa sahip text bloklarının yatay kaydırma durumunu kontrol eder
document.querySelectorAll('.text-block').forEach(block => {
  if (block.scrollWidth > block.clientWidth) {
    block.style.overflow = 'auto';
  }
});

// Karakter sayacı ekler
var textarea = document.getElementById("content");
var charCount = document.getElementById("charCount");

textarea.addEventListener("input", function () {
  var charLength = textarea.value.length;
  charCount.textContent = "Characters: " + charLength + "/600";

  if (charLength > 600) {
    charCount.style.color = "red";
  } else {
    charCount.style.color = "black";
  }
});
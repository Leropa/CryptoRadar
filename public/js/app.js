// Адрес нашего Go-сервера
const API_URL = 'http://localhost:8080/api/arbitrage';

// Функция получения данных с бэкенда
async function fetchCryptoData() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error(`Ошибка сервера: ${response.status}`);
        }
        const data = await response.json();
        renderCards(data);
    } catch (error) {
        console.error('Ошибка подключения к бэкенду:', error);
        showError();
    }
}

// Функция отрисовки карточек на странице
function renderCards(coins) {
    const container = document.getElementById('cards-container');
    container.innerHTML = ''; // Очищаем контейнер от старых данных

    coins.forEach(coin => {
        const card = document.createElement('div');
        // Применяем стили Tailwind и наш кастомный класс .crypto-card
        card.className = 'crypto-card border border-slate-800 rounded-2xl p-6 shadow-xl';

        card.innerHTML = `
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-xl font-bold tracking-tight">${coin.name}</h2>
                <span class="text-xs bg-slate-700 text-slate-300 px-2.5 py-1 rounded-md font-mono">${coin.symbol}</span>
            </div>
            
            <div class="space-y-3 border-b border-slate-700/50 pb-4 mb-4 text-sm">
                <div class="flex justify-between">
                    <span class="text-slate-400">Binance:</span>
                    <span class="font-mono font-semibold">$${coin.binance_price.toLocaleString()}</span>
                </div>
                <div class="flex justify-between">
                    <span class="text-slate-400">Bybit:</span>
                    <span class="font-mono font-semibold">$${coin.bybit_price.toLocaleString()}</span>
                </div>
            </div>

            <div class="bg-slate-900/60 p-3 rounded-xl flex justify-between items-center">
                <div>
                    <p class="text-[10px] text-slate-500 uppercase font-bold tracking-wider">Лучшая сделка</p>
                    <p class="text-xs text-emerald-400 font-medium mt-0.5">Купить на ${coin.best_buy_at}</p>
                </div>
                <div class="text-right">
                    <span class="text-[10px] text-slate-500 uppercase font-bold tracking-wider block">Спред</span>
                    <span class="text-sm font-mono font-bold text-cyan-400">+${coin.spread.toFixed(2)}%</span>
                </div>
            </div>
        `;
        container.appendChild(card);
    });
}

// Функция показа ошибки, если бэк не отвечает
function showError() {
    const container = document.getElementById('cards-container');
    container.innerHTML = `
        <div class="col-span-3 bg-red-950/30 border border-red-800 text-red-400 p-4 rounded-xl text-center">
            Ошибка подключения к бэкенду на <code class="bg-black/30 px-1 rounded">localhost:8080</code>. Убедись, что твой сервер на Go запущен.
        </div>
    `;
}

// Запускаем опрос сервера каждую секунду
setInterval(fetchCryptoData, 1000);

// Первый запуск при загрузке страницы
fetchCryptoData();
const socket = new WebSocket("ws://" + window.location.host + "/ws");

document.addEventListener("DOMContentLoaded", () => {
    loadNotifications();

    const bellIcon = document.getElementById("bell-icon");
    const dropdown = document.getElementById("notifications-dropdown");

    if (bellIcon && dropdown) {
        bellIcon.addEventListener("click", (e) => {
            e.stopPropagation();
            dropdown.style.display = dropdown.style.display === "block" ? "none" : "block";
        });

        document.addEventListener("click", () => {
            dropdown.style.display = "none";
        });

        dropdown.addEventListener("click", (e) => e.stopPropagation());
    }
});

// Загрузка истории уведомлений
async function loadNotifications() {
    try {
        const res = await fetch("/api/notifications/list");
        if (!res.ok) return;
        
        const data = await res.json();

        const listContainer = document.getElementById("notifications-list");
        if (!listContainer) return;

        if (!data.items || data.items.length === 0) {
            listContainer.innerHTML = '<div class="empty-notifications">Нет новых уведомлений</div>';
            return;
        }

        listContainer.innerHTML = "";
        data.items.forEach(item => {
            appendNotificationDOM(item.id, item.text, item.is_read, false);
        });

    } catch (err) {
        console.error("Ошибка загрузки уведомлений:", err);
    }
}

// Добавление уведомления в выпадающий список
function appendNotificationDOM(id, text, isRead, prepend = true) {
    const listContainer = document.getElementById("notifications-list");
    if (!listContainer) return;

    if (listContainer.querySelector(".empty-notifications")) {
        listContainer.innerHTML = "";
    }

    const div = document.createElement("div");
    div.className = "notification-item";
    if (!isRead) {
        div.classList.add("unread");
    }
    div.innerText = text;

    // Клик просто снимает выделение и шлет запрос на бэкенд
    div.onclick = () => {
        readNotification(id, div);
    };

    if (prepend) {
        listContainer.insertBefore(div, listContainer.firstChild);
    } else {
        listContainer.appendChild(div);
    }
}

// Пометка уведомления как прочитанного
async function readNotification(id, element) {
    // Если оно уже прочитано (нет класса unread), игнорируем клик
    if (!element.classList.contains("unread")) return;

    // Если у уведомления есть ID из базы, отправляем запрос бэкенду
    if (id > 0) {
        try {
            const res = await fetch(`/api/notifications/read/${id}`, { method: "POST" });
            if (!res.ok) return;
        } catch (err) {
            console.error("Ошибка при чтении уведомления:", err);
            return;
        }
    }

    // Просто убираем класс unread (синее выделение пропадает)
    element.classList.remove("unread");
}

// Получение уведомлений по сокетам в реальном времени
socket.onmessage = function(event) {
    const notificationText = event.data;

    // Новое уведомление падает вверх списка, горит синим, id = 0
    appendNotificationDOM(0, notificationText, false, true);
};

socket.onclose = () => console.log("[WebSocket] Закрыто");
socket.onerror = (err) => console.error("[WebSocket] Ошибка:", err);
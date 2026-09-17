<script lang="ts">
    import { Quit } from "../wailsjs/runtime";
    import {
        GetAvailablePorts,
        ShowError,
        ShowMessage,
    } from "../wailsjs/go/main/App";
    import { onMount } from "svelte";

    let ports: string[] = [];
    let selectedPort: string = "";
    let password: string = "";

    // Функция для обновления списка портов
    async function refreshPorts() {
        try {
            ports = await GetAvailablePorts();
            // Если порты есть, а выбpанный порт пустой — выберем первый по умолчанию
            if (ports.length > 0 && !selectedPort) {
                selectedPort = ports[0];
            }
        } catch (error) {
            ShowError("Ошибка", "Не удалось получить порты");
        }
    }

    function runApp() {
        if (password.length > 0) {
            ShowMessage("Запуск программы ");
        } else {
            ShowError("Отказано в доступе", "Неправильный пароль пользователя");
        }
    }

    function runEmulator() {
        ShowMessage("Запуск эмулятора ");
    }

    function runConfig() {
        ShowMessage("Запуск конфигурации ");
    }

    onMount(() => {
        refreshPorts();
    });
</script>

<!-- ОТКЛЮЧАЕМ СТАНДАРТНОЕ КОНТЕКСТНОЕ МЕНЮ -->
<svelte:window on:contextmenu={(e) => e.preventDefault()} />

<main>
    <div class="header">
        <div class="title">Launcher</div>
        <div class="window-controls">
            <button class="wc-btn" id="minimizeApp" title="Выход" onclick={Quit}
            ></button>
        </div>
    </div>
    <div class="run-form">
        <div class="input-box">
            <label for="port">Порт</label>
            <select name="port" id="port">
                {#if ports.length === 0}
                    <option value="">Порты не найдены</option>
                {:else}
                    {#each ports as port}
                        <option value={port}>{port}</option>
                    {/each}
                {/if}
            </select>
        </div>
        <div class="input-box">
            <label for="user">Пользователь</label>
            <select name="user" id="user">
                <option value="user1">Оператор 1</option>
            </select>
        </div>
        <div class="input-box">
            <label for="password">Пароль</label>
            <input
                type="password"
                name="password"
                id="password"
                value={password}
            />
        </div>
        <div class="spacer"></div>
        <div class="input-box">
            <button onclick={runApp}>Старт</button>
            <button onclick={runEmulator}>Имитатор УСО</button>
            <button onclick={runConfig}>Конфигурация</button>
        </div>
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        user-select: none;
    }
    .run-form {
        padding: 32px;
    }
    .spacer {
        height: 20px;
    }

    .input-box {
        display: flex;
        align-items: center;
        margin-bottom: 8px;
        gap: 8px;

        select,
        input,
        button {
            font-family: "internal-app-font";
            font-size: var(--font-size);
            border-radius: 5px;
            flex: 1;
            appearance: none;
            border: 0;
            cursor: pointer;
            color: var(--text-color);
            background-color: var(--inputs-background-color);
            padding-left: 16px;
            height: 38px;
            border-bottom: 1px solid var(--border-color);
        }

        select,
        input {
            &:focus {
                outline: 0;
                background-color: #252525;
                color: var(--primary-color);
            }
        }

        select option {
            font-size: var(--font-size);
            background-color: #222222;
            color: var(--text-color);
        }

        label {
            flex: 1;
            font-size: var(--font-size);
            color: var(--text-color);
            text-align: right;
        }

        button {
            padding: 0;
            height: 40px;
            color: var(--text-color);
            background-color: var(--button-background-color);

            box-shadow:
                0 -1px 1px rgba(255, 255, 255, 0.45),
                0 4px 4px rgba(0, 0, 0, 0.1),
                0 19px 15px rgba(255, 255, 255, 0.14) inset;

            border: none;
            &:hover {
                background-color: var(--button-background-color-hover);
                box-shadow:
                    0 -1px 1px rgba(255, 255, 255, 0.65),
                    0 19px 15px rgba(255, 255, 255, 0.2) inset;
            }
            &:active {
                padding-top: 2px;
                background-color: var(--button-background-color-active);
                box-shadow:
                    0 1px 1px rgba(255, 255, 255, 0.25),
                    0 -1px 1px rgba(0, 0, 0, 0.1),
                    0 19px 15px rgba(255, 255, 255, 0.1) inset;
            }
        }
    }
</style>

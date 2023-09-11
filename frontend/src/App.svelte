<script>
  import { Greet } from "../wailsjs/go/main/App.js";
  import { Register } from "../wailsjs/go/main/App.js";
  import { LoadAll } from "../wailsjs/go/main/App.js";
  import { main } from "../wailsjs/go/models";

  let resultText = "Please enter your name below 👇";
  let regiterResult = false;
  let name;
  const done = load();

  function greet() {
    Greet(name).then((result) => (resultText = result));
  }

  function register() {
    let e = new main.Entry();
    e.type = "Process";
    Register(e).then((result) => (regiterResult = result));
  }

  async function load() {
    console.log("sart loading");
    const jobs = await LoadAll();
    console.log("loaded");
    return jobs;
  }
</script>

<main>
  <div class="result" id="result">{resultText}</div>
  <div class="result" id="regiterResult">{regiterResult}</div>
  <div class="input-box" id="input">
    <input
      autocomplete="off"
      bind:value={name}
      class="input"
      id="name"
      type="text"
    />
    <button class="btn" on:click={greet}>Greet</button>
  </div>
  <button class="btn" on:click={register}>Register</button>
  <div class="grid">
    {#await done}
      Loading1
    {:then jarray}
      {#each jarray as j (j.handle)}
        <button class="btn">handle is {j.handle}</button>
      {/each}
    {/await}
  </div>
</main>

<style>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>

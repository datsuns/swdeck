<script>
  import { Add, Get, Launch, Delete } from "../wailsjs/go/main/App.js";
  //import { Job } from "../wailsjs/go/models";
  import { onMount } from "svelte";

  let jobs = [];
  let jobName = "";
  let jobCommand = "";
  let jobType = "app"; // デフォルトとして 'app' を選択

  onMount(async () => {
    jobs = await Get();
  });

  async function addJob() {
    if (jobName && jobCommand) {
      const newJob = { name: jobName, command: jobCommand, type: jobType };
      await Add(newJob);
      jobs = await Get();
      jobName = "";
      jobCommand = "";
    }
  }

  async function deleteJob(index) {
    await Delete(index);
    jobs = await Get();
  }

  function launchJob(job) {
    Launch(job);
  }
</script>

<input bind:value={jobName} placeholder="Job Name" />
<input bind:value={jobCommand} placeholder="Command or Path" />

<!-- Job Type の選択 -->
<select bind:value={jobType}>
  <option value="app">Application</option>
  <option value="mp3">MP3 File</option>
</select>

<button on:click={addJob}>Add Job</button>

<ul>
  {#each jobs as job, index}
    <li>
      <button on:click={() => launchJob(job)}>{job.name}</button>
      <button on:click={() => deleteJob(index)}>Delete</button>
    </li>
  {/each}
</ul>

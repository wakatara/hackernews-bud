<script>
  import Header from './Header.svelte'
  import Story from './Story.svelte'
  import Comment from './Comment.svelte'
  export let story = {}
</script>

<Header />

<Story {story} />

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(story) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(story) as key}
      <td>{story[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href="">Back</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>

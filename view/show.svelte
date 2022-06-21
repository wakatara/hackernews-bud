<script>
  import Header from './Header.svelte'
  import Story from './Story.svelte'
  import Comment from './Comment.svelte'
  export let story = {}
</script>

<Header />

<Story {story} />

{#each story.children as comment}
  <Comment {comment} />
{/each}


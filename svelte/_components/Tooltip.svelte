<script>
  let x = 0;
	let y = 0;
  let isHovered = false;

  /**
   * @param {DOMRect} rect
   */
  export function Show(rect) {
    isHovered = true;

		x = rect.left + 10;
		y = rect.top + 10;
  }

  export function Hide() {
    isHovered = false;
  }

</script>

{#if isHovered}
  <div style="top: {y}px; left: {x}px;" class="tooltip" >
    <slot />
  </div>
{/if}

<style>
	.tooltip {
		border: 2px solid var(--gray-003);
		box-shadow: 1px 1px 1px var(--gray-006);
		background: #FFF;
		border-radius: 5px;
		padding: 10px;
		position: absolute;
    display: flex;
    flex-direction: column;
    gap: 5px;
	}
</style>
<script>
  import { tick } from "svelte";

  let tooltip = /** @type {HTMLDivElement} */ (null);

  let x = 0;
	let y = 0;
  let isHovered = false;

  /**
   * @param {DOMRect} rect
   */
  export async function Show(rect) {
    isHovered = true;

    await tick();

    const tooltipRect = tooltip.getBoundingClientRect();

    const windowHeight = window.innerHeight;
    const windowWidth = window.innerWidth;


    x = rect.left + 30;
    y = rect.top + 30;

    if (y + tooltipRect.height > windowHeight) {
      y = rect.top - tooltipRect.height - 30;
    }

    if (x + tooltipRect.width > windowWidth) {
      x = rect.left - tooltipRect.width - 30;
    }
  }

  export function Hide() {
    isHovered = false;
  }
</script>

<div
  bind:this={tooltip}
  style="top: {y}px; left: {x}px;"
  class="tooltip"
  class:show={isHovered}
>
  <slot />
</div>

<style>
	.tooltip {
		border: 2px solid var(--gray-003);
		box-shadow: 1px 1px 1px var(--gray-006);
		background-color: #FFF;
		border-radius: 5px;
		padding: 10px;
		position: fixed;
    z-index: 101;
    display: none;
    flex-direction: column;
    gap: 5px;
	}

  .tooltip.show {
    display: flex;
  }
</style>
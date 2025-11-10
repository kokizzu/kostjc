<script>
	export let title = '';

	let isHovered = false;
	let x;
	let y;
  const offset = 10; // Gap between button and tooltip

  let tooltip = /** @type {HTMLDivElement} */ (null);

  function positionTooltip(btn) {
      const btnRect = btn.getBoundingClientRect();
      const tooltipRect = tooltip.getBoundingClientRect();
      const viewportWidth = window.innerWidth;
      const viewportHeight = window.innerHeight;
      
      let top, left;
      let position = 'top'; // default position
      
      // Calculate if tooltip fits on top
      const fitsTop = btnRect.top - tooltipRect.height - offset > 0;
      const fitsBottom = btnRect.bottom + tooltipRect.height + offset < viewportHeight;
      const fitsLeft = btnRect.left - tooltipRect.width - offset > 0;
      const fitsRight = btnRect.right + tooltipRect.width + offset < viewportWidth;
      
      // Remove all position classes
      tooltip.className = 'tooltip show';
      
      // Determine best position (priority: top, bottom, left, right)
      if (fitsTop) {
        // Position above
        position = 'top';
        top = btnRect.top - tooltipRect.height - offset;
        left = btnRect.left + (btnRect.width / 2) - (tooltipRect.width / 2);
      } else if (fitsBottom) {
        // Position below
        position = 'bottom';
        top = btnRect.bottom + offset;
        left = btnRect.left + (btnRect.width / 2) - (tooltipRect.width / 2);
      } else if (fitsRight) {
        // Position to the right
        position = 'right';
        top = btnRect.top + (btnRect.height / 2) - (tooltipRect.height / 2);
        left = btnRect.right + offset;
      } else if (fitsLeft) {
        // Position to the left
        position = 'left';
        top = btnRect.top + (btnRect.height / 2) - (tooltipRect.height / 2);
        left = btnRect.left - tooltipRect.width - offset;
      } else {
        // Fallback: position below (even if it doesn't fit perfectly)
        position = 'bottom';
        top = btnRect.bottom + offset;
        left = btnRect.left + (btnRect.width / 2) - (tooltipRect.width / 2);
      }
      
      // Adjust horizontal position if tooltip would overflow viewport
      if (position === 'top' || position === 'bottom') {
        if (left < offset) {
          left = offset;
        } else if (left + tooltipRect.width > viewportWidth - offset) {
          left = viewportWidth - tooltipRect.width - offset;
        }
      }
      
      // Adjust vertical position if tooltip would overflow viewport
      if (position === 'left' || position === 'right') {
        if (top < offset) {
          top = offset;
        } else if (top + tooltipRect.height > viewportHeight - offset) {
          top = viewportHeight - tooltipRect.height - offset;
        }
      }
      
      // Apply position
      tooltip.style.top = `${top}px`;
      tooltip.style.left = `${left}px`;
      tooltip.classList.add(position);
  }
	
	function mouseOver(event) {
		isHovered = true;
		x = event.pageX + 5;
		y = event.pageY + 5;
	}
	function mouseMove(event) {
		x = event.pageX + 5;
		y = event.pageY + 5;
	}
	function mouseLeave() {
		isHovered = false;
	}

  function onfocus() {

  }

  function onblur() {

  }
</script>

<div
  bind:this={tooltip}
  aria-label="tooltip"
  role="tooltip"
	on:mouseover={mouseOver}
  on:mouseleave={mouseLeave}
	on:mousemove={mouseMove}
  on:focus={onfocus}
  on:blur={onblur}
>
	<slot />
</div>

{#if isHovered}
	<div style="top: {y}px; left: {x}px;" class="tooltip">{title}</div>
{/if}

<style>
	.tooltip {
		border: 1px solid #ddd;
		box-shadow: 1px 1px 1px #ddd;
		background: white;
		border-radius: 4px;
		padding: 4px;
		position: absolute;
	}
</style>
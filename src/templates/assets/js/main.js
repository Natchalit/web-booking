let scrollPosition = 0;

function scrollBlog(direction) {
  const blogContainer = document.getElementById("blogContainer");
  const blogItemWidth = document.querySelector(".blog-item").offsetWidth;

  // Calculate the new scroll position based on the direction
  scrollPosition += direction * (blogItemWidth + 2); // 2 for the column-gap

  // Limit the scroll position to prevent scrolling beyond the content
  scrollPosition = Math.max(
    0,
    Math.min(
      scrollPosition,
      blogContainer.scrollWidth - blogContainer.clientWidth
    )
  );

  // Apply the new scroll position
  blogContainer.scrollTo({ left: scrollPosition, behavior: "smooth" });
}

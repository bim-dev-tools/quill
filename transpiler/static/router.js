const loadContent = async (page) => {
  try {
    const response = await fetch(`posts/${page}.html`);
    const text = await response.text();

    // Extract only the main content
    const parser = new DOMParser();
    const doc = parser.parseFromString(text, "text/html");
    const newContent = doc.body;

    return newContent;
  } catch (error) {
    console.error("Error loading content:", error);
  }
};

const renderContent = async (path, push = true) => {
  const pageName = path.match(/(?<=.*#\/).+$/g);

  const pageContent = await loadContent(pageName || "_home");

  if (pageContent) {
    document.querySelector("#spa-content").replaceChildren(pageContent);
    const url = pageName ? `/#/${pageName}` : "/";
    push ? history.pushState(url, "", url) : history.replaceState(url, "", url); // Update browser history

    Prism.highlightAll();
  }
};

const interceptRouting = async (event) => {
  event.preventDefault();

  await renderContent(event.target.href);
};

document.addEventListener("click", async (e) => {
  if (e.target.dataset["nav"] === "true") {
    e.preventDefault();

    e.target.href !== location.href && (await renderContent(e.target.href));
  }
});

window.addEventListener("popstate", async (_event) => {
  await renderContent(location.hash, false);
});

(async () => {
  await renderContent(location.hash, false);
})();

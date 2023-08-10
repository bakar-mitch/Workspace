const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

const namesArray = [
  "Alice", "Bob", "Charlie", "David", "Emma",
  "Frank", "Grace", "Henry", "Ivy", "Jack",
  "Kate", "Leo", "Mia", "Nora", "Oliver",
  "Penny", "Quincy", "Riley", "Sophia", "Thomas",
  "Uma", "Victoria", "William", "Xander", "Yara", "Zoe"
];

function visualizeBinarySearch(target) {
  let left = 0;
  let right = namesArray.length - 1;

  while (left <= right) {
    const middle = Math.floor((left + right) / 2);
    // Visualization code here (similar to previous example)

    if (namesArray[middle] === target) {
      showResultCard(`Found ${target} at index ${middle}`);
      return;
    } else if (namesArray[middle] < target) {
      left = middle + 1;
    } else {
      right = middle - 1;
    }
  }

  showResultCard(`${target} not found in the array`);
}

function showResultCard(message) {
  const resultCard = document.getElementById("resultCard");
  const resultText = document.getElementById("resultText");
  resultText.textContent = message;
  resultCard.style.display = "block";
}

const searchButton = document.getElementById("searchButton");
searchButton.addEventListener("click", () => {
  const target = document.getElementById("target").value;
  visualizeBinarySearch(target);
});

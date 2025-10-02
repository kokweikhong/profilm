document.addEventListener("DOMContentLoaded", function () {
  warrantyTypeSelect = document.getElementById("warranty_type");
  warrantyContentDiv = document.getElementById("warranty_content");
  ppfWarrantyContent = document.getElementById("ppf_warranty_content");
  windowFilmWarrantyContent = document.getElementById(
    "window_film_warranty_content"
  );
  sunroofIceShieldWarrantyContent = document.getElementById(
    "sunroof_ice_shield_warranty_content"
  );

  warrantyTypeSelect.addEventListener("change", function () {
    var selectedValue = this.value;

    // Hide all content divs
    warrantyContentDiv.style.display = "none";
    ppfWarrantyContent.style.display = "none";
    windowFilmWarrantyContent.style.display = "none";
    sunroofIceShieldWarrantyContent.style.display = "none";

    // Show the selected content div
    if (selectedValue === "ppf") {
      ppfWarrantyContent.style.display = "block";
      warrantyContentDiv.style.display = "block";
    } else if (selectedValue === "window_film") {
      windowFilmWarrantyContent.style.display = "block";
      warrantyContentDiv.style.display = "block";
    } else if (selectedValue === "sunroof_ice_shield") {
      sunroofIceShieldWarrantyContent.style.display = "block";
      warrantyContentDiv.style.display = "block";
    }
  });
});

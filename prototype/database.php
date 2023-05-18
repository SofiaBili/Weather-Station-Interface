<?php
// Assuming you have already established a database connection

// Accessing values from the bme280_values table
$query = "SELECT * FROM bme280_values";
$result = mysqli_query($connection, $query);

if ($result) {
    while ($row = mysqli_fetch_assoc($result)) {
        $sensorId = $row['sensor_id'];
        $moment = $row['moment'];
        $temperature = $row['temperature'];
        $humidity = $row['humidity'];
        $barometricPressure = $row['barometric_pressure'];

        // Do something with the values
        // ...
    }
}else{
        // Handle query error
        echo "Error: " . mysqli_error($connection);
}

// Accessing values from the dust_sensor_values table
$query = "SELECT * FROM dust_sensor_values";
$result = mysqli_query($connection, $query);
while ($row = mysqli_fetch_assoc($result)) {
    $sensorId = $row['sensor_id'];
    $moment = $row['moment'];
    $dustDensity = $row['dust_density'];

    // Do something with the values
    // ...
}

// Accessing values from the mq2_values table
$query = "SELECT * FROM mq2_values";
$result = mysqli_query($connection, $query);
while ($row = mysqli_fetch_assoc($result)) {
    $sensorId = $row['sensor_id'];
    $moment = $row['moment'];
    $h2 = $row['h2'];
    $lpg = $row['lpg'];
    $propane = $row['propane'];

    // Do something with the values
    // ...
}

// Accessing values from the mq3_values table
$query = "SELECT * FROM mq3_values";
$result = mysqli_query($connection, $query);
while ($row = mysqli_fetch_assoc($result)) {
    $sensorId = $row['sensor_id'];
    $moment = $row['moment'];
    $alcohol = $row['alcohol'];
    $benzene = $row['benzene'];
    $hexane = $row['hexane'];

    // Do something with the values
    // ...
}

// Accessing values from the mq4_values table
$query = "SELECT * FROM mq4_values";
$result = mysqli_query($connection, $query);
while ($row = mysqli_fetch_assoc($result)) {
    $sensorId = $row['sensor_id'];
    $moment = $row['moment'];
    $ch4 = $row['ch4'];
    $co = $row['co'];
    $smoke = $row['smoke'];

    // Do something with the values
    // ...
}

// Accessing values from the mq135_values table
$query = "SELECT * FROM mq135_values";
$result = mysqli_query($connection, $query);
while ($row = mysqli_fetch_assoc($result)) {
    $sensorId = $row['sensor_id'];
    $moment = $row['moment'];
    $co2 = $row['co2'];
    $tolueno = $row['tolueno'];
    $nh4 = $row['nh4'];
    $acetone = $row['acetone'];

    // Do something with the values
    // ...
}

// Accessing values from the rain_sensor



// Close the database connection
mysqli_close($connection);
?>
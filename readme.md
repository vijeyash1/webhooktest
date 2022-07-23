https://f973-122-178-160-70.in.ngrok.iohttps://f973-122-178-160-70.in.ngrok.iohttps://f973-122-178-160-70.in.ngrok.io
https://f973-122-178-160-70.in.ngrok.io
https://f973-122-178-160-70.in.ngrok.io
https://f973-122-178-160-70.in.ngrok.iohttps://f973-122-178-160-70.in.ngrok.io
https://meet.google.com/ndv-ugkm-yxo
https://meet.google.com/ndv-ugkm-yxo

https://630b-122-178-160-70.in.ngrok.iohttps://630b-122-178-160-70.in.ngrok.io

    db.AutoMigrate(&models.Gitevent{})
    // Set table options
    db.Set("gorm:table_options", "ENGINE=File(cluster, default, hits)").AutoMigrate(&models.Gitevent{})

    // Set table cluster options
    db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&models.Gitevent{})

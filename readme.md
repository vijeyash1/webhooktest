this a a github test payload

this a a github test payload

this a a github test payload
this a a github test payload
this a a github test payload
this a a github test payload
this is a gitlab webhook
this a a github test payload
this a a github test payload
this a a github test payload
this a a github test payload
this a a github test payloadthis this a a github test payload
a a github test payload
this a a github test payload
this a a github test payload
73c6a473c6a473c6a473c6a473c6a4
this a a github test payload73c6a473c6a473c6a473c6a473c6a473c6a4

this a a github test payload
this is a gitlab webhook

this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhook

this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhook

this is a gitlab webhook

this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook

this is a gitlab webhook

this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
this is a gitlab webhook

this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
this is a gitlab webhook

this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
is is a gitlab webhookthis is a gitlab webhook
this is a gitlab webhook
this is a gitlab webhookthis is a gitlab webhook
is a gitlab webhook

git add . && git commit -m "ok" && git pushgit add . && git commit -m "ok" && git pushgit add . && git commit -m "ok" && git pushgit add . && git commit -m "ok" && git push

git add . && git commit -m "ok" && git pushgit add . && git commit -m "ok" && git pushgit add . && git commit -m "ok" && git push
this is a gitlab webhookthis is a gitlab webhook

db, err := gorm.Open(clickhouse.Open(url), &gorm.Config{})
if err != nil {
panic("failed to connect database")
}
// Auto Migrate
db.AutoMigrate(&models.Gitevent{})
// Set table options
db.Set("gorm:table_options", "ENGINE=File(cluster, default, hits)").AutoMigrate(&models.Gitevent{})

    // Set table cluster options
    db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&models.Gitevent{})

    // Insert
    db.Create(&models.Gitevent{Uuid: string(metrics.Uuid), Event: metrics.Event, Eventid: metrics.Eventid, Branch: metrics.Branch, Url: metrics.Url, Authorname: metrics.Authorname, Authormail: metrics.Authormail, DoneAt: metrics.DoneAt, Repository: metrics.Repository, Addedfiles: metrics.Addedfiles, Modifiedfiles: metrics.Modifiedfiles, Removedfiles: metrics.Removedfiles, Message: metrics.Message})db, err := gorm.Open(clickhouse.Open(url), &gorm.Config{})
    if err != nil {
    	panic("failed to connect database")
    }
    // Auto Migrate
    db.AutoMigrate(&models.Gitevent{})
    // Set table options
    db.Set("gorm:table_options", "ENGINE=File(cluster, default, hits)").AutoMigrate(&models.Gitevent{})

    // Set table cluster options
    db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&models.Gitevent{})

    // Insert
    db.Create(&models.Gitevent{Uuid: string(metrics.Uuid), Event: metrics.Event, Eventid: metrics.Eventid, Branch: metrics.Branch, Url: metrics.Url, Authorname: metrics.Authorname, Authormail: metrics.Authormail, DoneAt: metrics.DoneAt, Repository: metrics.Repository, Addedfiles: metrics.Addedfiles, Modifiedfiles: metrics.Modifiedfiles, Removedfiles: metrics.Removedfiles, Message: metrics.Message})

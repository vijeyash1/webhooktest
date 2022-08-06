select event.metrics.after as hashid,
event.metrics.repository.name as repo,
event.metrics.commits.committer.name as committer,
event.metrics.commits.message as message,
event.metrics.commits.timestamp as time
from test.git_json where repo =='Gitea' SETTINGS describe_extend_object_types=1;
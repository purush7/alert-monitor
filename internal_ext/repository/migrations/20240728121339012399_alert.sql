CREATE TABLE alert (
    id SERIAL PRIMARY KEY,
    client VARCHAR(255) NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    alert_config_type VARCHAR(255) NOT NULL,
    count INT NOT NULL,
    window_size_in_secs INT NOT NULL,
    dispatcher_config JSON
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    client VARCHAR(255) NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

create index  event_created_at on events(created_at);
create index  alert_client_event_type on alert(client,event_type);
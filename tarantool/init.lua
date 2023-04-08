box.cfg { listen = 3301 }

s = box.schema.space.create('test_space_vinyl', {
    engine = 'vinyl',
    if_not_exists = true
})

s:format({
    {
        name = 'domain',
        type = 'string'
    },
    {
        name = 'type',
        type = 'string'
    },
    {
        name = "click",
        type = "unsigned",
    }
})

s:create_index('primary', {
    if_not_exists = true,
    unique = true,
    type = 'TREE',
    parts = {
        'domain',
        'type',
    }
})

s:create_index('domain', {
    if_not_exists = true,
    unique = false,
})

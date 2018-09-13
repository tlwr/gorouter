# 2. Change TLS route pruning behavior on dial error

Date: 2018-09-12

## Status

Accepted

## Context

This is related to story [#158847588](https://www.pivotaltracker.com/story/show/158847588)

In the story above, when route-integrity is turned on
(soon the default) we do not prune routes that receive a `Dial` error.
The code has special logic in place making
it so that there are only two types of errors that result in a
prune:

[errors on which we can prune](https://github.com/cloudfoundry/gorouter/blob/b63e6fb16c2a422ec5108a19debc9adb81f2d1dd/proxy/fails/classifier_group.go#L17-L20)

The prune operation should be harmless - route will get added again
(if the application is still around) the next time the
route-registrar runs.

[current code behavior with classifier](https://github.com/cloudfoundry/gorouter/blob/b63e6fb16c2a422ec5108a19debc9adb81f2d1dd/route/pool.go#L369-L372)

## Decision

We are going to prune a TLS route on a `Dial` error _if_
the TTL of the route has passed. This gives the
route time to recover.

We will also be logging the cases where we are pruning versus
marking an endpoint as failed and removing it from the route
pool for a 'cooldown'.

## Consequences

If developers app is not coming up they will see the same error
pattern of 502s they saw before, the current behavior does not
change.

Logging will be introduced into the route pool for the first time
giving operators a view into whether a prune or fail has actually
occurred and what error caused it.

Now TLS routes will actually get removed from the pool if there
is an error and their TTL has passed.

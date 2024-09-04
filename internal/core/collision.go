package core

type CollisionArea interface {
	OnHit(collider any)
	Collision
}

type Collider interface {
	OnCollided(collider any)
	Collision
}

type Collision interface {
	// GetMetadataForCollision should return the entity x, y, width, height
	GetMetadataForCollision() (int32, int32, int32, int32)
}

func NewCollisionServer() *CollisionServer {
	collisionServer := CollisionServer{}

	return &collisionServer
}

// NOTE: This worked!
type CollisionServer struct {
	CollisionAreas []*CollisionArea
	Colliders      []*Collider
}

func (c *CollisionServer) RegisterNode(node any) {
	if collisionArea, ok := node.(CollisionArea); ok {
		c.CollisionAreas = append(c.CollisionAreas, &collisionArea)
	}
	if collider, ok := node.(Collider); ok {
		c.Colliders = append(c.Colliders, &collider)
	}
}

func (c *CollisionServer) checkCollision(x1, y1, w1, h1, x2, y2, w2, h2 int32) bool {
	return (max(x1, x2) < min(x1+w1, x2+w2)) && (max(y1, y2) < min(y1+h1, y2+h2))
}

func (c *CollisionServer) Scan() {
	for _, collider := range c.Colliders {
		x1, y1, w1, h1 := (*collider).GetMetadataForCollision()

		for _, collisionArea := range c.CollisionAreas {
			x2, y2, w2, h2 := (*collisionArea).GetMetadataForCollision()

			if c.checkCollision(x1, y1, w1, h1, x2, y2, w2, h2) {
				(*collider).OnCollided(collisionArea)
				(*collisionArea).OnHit(collider)
			}
		}
	}
}

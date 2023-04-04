import EntityList = models.EntityList
import RelatedEntityList = models.RelatedEntityList
import EntityItem = models.EntityItem

export async function searchEntities(search: string): Promise<EntityList> {
	const res = await fetch(`/api/v2/entities_search/${search}`)
	if (!res.ok) {
		throw new Error(`failed to get feed: ${res.status}`)
	}
	return res.json()
}

export async function getEntities(): Promise<EntityList> {
	const res = await fetch(`/api/v2/entities`)
	if (!res.ok) {
		throw new Error(`failed to get feed: ${res.status}`)
	}
	return res.json()
}

export async function loadRelatedEntities(items: EntityItem[][][]): Promise<RelatedEntityList> {
	const res = await fetch(`/api/v2/entities_involved`, {
		method: "POST",
		body: JSON.stringify({ items }),
		headers: {
			"Content-Type": "application/json",
		},
	})
	if (!res.ok) {
		throw new Error(`failed to get related entities: ${res.status}`)
	}
	return res.json()
}

export function filterEntities(
	srcStr: string,
	list: EntityList,
	trackItems: EntityItem[][][],
	ignoreItems: EntityItem[]
): EntityList {
	srcStr = srcStr.toLowerCase()
	const result: EntityList = {}
	// filter all items in list and return result match
	for (let entityType in list) {
		result[entityType] = []
		for (let j in list[entityType]) {
			if (!list[entityType][j].src_name.toLowerCase().includes(srcStr)) {
				continue
			}
			if (
				// todo disabled, while allow use same item multuply time checkItemExistInList(list[entityType][j], trackItems) ||
				checkItemExist(list[entityType][j], ignoreItems)
			) {
				continue
			}
			result[entityType].push(list[entityType][j])
		}
		result[entityType].sort((a, b) => b.count - a.count)
	}
	return result
}

export function filterRelatedEntities(
	relatedItems: RelatedEntityList,
	trackItems: EntityItem[][][],
	ignoreItems: EntityItem[]
): RelatedEntityList {
	let result: RelatedEntityList = {}
	for (let entityType in relatedItems) {
		result[entityType] = {}
		for (let relatedEntityType in relatedItems[entityType]) {
			result[entityType][relatedEntityType] = []
			for (let i in relatedItems[entityType][relatedEntityType]) {
				if (
					checkItemExistInList(relatedItems[entityType][relatedEntityType][i], trackItems) ||
					checkItemExist(relatedItems[entityType][relatedEntityType][i], ignoreItems)
				) {
					continue
				}
				result[entityType][relatedEntityType].push(relatedItems[entityType][relatedEntityType][i])
			}
			result[entityType][relatedEntityType].sort((a, b) => b.count - a.count)
		}
	}
	return result
}

function checkItemExist(item: EntityItem, list: EntityItem[]): boolean {
	for (let i in list) {
		if (list[i].src_name === item.src_name && list[i].entity_type === item.entity_type) {
			return true
		}
	}
	return false
}

function checkItemExistInList(item: EntityItem, list: EntityItem[][][]): boolean {
	for (let i in list) {
		for (let j in list[i]) {
			if (checkItemExist(item, list[i][j])) {
				return true
			}
		}
	}
	return false
}

#ifndef MARISA_GLUE_H_
#define MARISA_GLUE_H_

#include "marisa/base.h"
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void *marisa_keyset;
typedef void *marisa_trie;
typedef void *marisa_agent;

typedef struct {
    marisa_error_code err;
    bool found;
    size_t id;
    const char *str;
    size_t len;
} marisa_result;

typedef struct {
    marisa_error_code err;
    size_t num;
} marisa_num;


marisa_keyset marisa_keyset_new();
void marisa_keyset_delete(marisa_keyset keyset);

marisa_error_code marisa_keyset_push_back(marisa_keyset keyset, const char *ptr, size_t length, float weight);


marisa_agent marisa_agent_new();
void marisa_agent_delete(marisa_agent agent);

marisa_error_code marisa_agent_set_query_with_str(marisa_agent agent, const char* str, size_t len);
marisa_error_code marisa_agent_set_query_with_id(marisa_agent agent, size_t id);


marisa_trie marisa_trie_new();
void marisa_trie_delete(marisa_trie trie);

marisa_error_code marisa_trie_build(marisa_trie trie, const marisa_keyset keyset, int config_flags);
marisa_error_code marisa_trie_load(marisa_trie trie, const char* file);
marisa_error_code marisa_trie_mmap(marisa_trie trie, const char* file);
marisa_error_code marisa_trie_save(const marisa_trie trie, const char* file);

marisa_result marisa_trie_lookup(const marisa_trie trie, marisa_agent agent);
marisa_result marisa_trie_reverse_lookup(const marisa_trie trie, marisa_agent agent);
marisa_result marisa_trie_common_prefix_search(const marisa_trie trie, marisa_agent agent);
marisa_result marisa_trie_predictive_search(const marisa_trie trie, marisa_agent agent);

marisa_num marisa_trie_num_tries(const marisa_trie trie);
marisa_num marisa_trie_num_keys(const marisa_trie trie);
marisa_num marisa_trie_num_nodes(const marisa_trie trie);

marisa_num marisa_trie_size(const marisa_trie trie);
marisa_num marisa_trie_total_size(const marisa_trie trie);
marisa_num marisa_trie_io_size(const marisa_trie trie);

#ifdef __cplusplus
}
#endif
#endif
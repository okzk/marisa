
#include "marisa_glue.h"
#include "marisa/trie.h"

marisa_keyset marisa_keyset_new() {
    return new marisa::Keyset();
}

void marisa_keyset_delete(marisa_keyset keyset) {
    delete static_cast<marisa::Keyset*>(keyset);
}

marisa_error_code marisa_keyset_push_back(marisa_keyset keyset, const char *ptr, size_t length, float weight) {
    try {
        static_cast<marisa::Keyset*>(keyset)->push_back(ptr, length, weight);
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}


marisa_agent marisa_agent_new() {
    return new marisa::Agent();
}

void marisa_agent_delete(marisa_agent agent) {
    delete static_cast<marisa::Agent*>(agent);
}

marisa_error_code marisa_agent_set_query_with_str(marisa_agent agent, const char* str, size_t len) {
    marisa::Agent* ma = static_cast<marisa::Agent*>(agent);
    try {
        ma->set_query(str, len);
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}

marisa_error_code marisa_agent_set_query_with_id(marisa_agent agent, size_t id) {
    marisa::Agent* ma = static_cast<marisa::Agent*>(agent);
    try {
        ma->set_query(id);
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}


marisa_trie marisa_trie_new() {
    return new marisa::Trie();
}

void marisa_trie_delete(marisa_trie trie) {
    delete static_cast<marisa::Trie*>(trie);
}

marisa_error_code marisa_trie_build(marisa_trie trie, const marisa_keyset keyset, int config_flags) {
    try {
        static_cast<marisa::Trie*>(trie)->build(*static_cast<marisa::Keyset*>(keyset));
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}

marisa_error_code marisa_trie_load(marisa_trie trie, const char* file) {
    try {
        static_cast<marisa::Trie*>(trie)->load(file);
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}

marisa_error_code marisa_trie_mmap(marisa_trie trie, const char* file) {
    try {
        static_cast<marisa::Trie*>(trie)->mmap(file);
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}


marisa_error_code marisa_trie_save(const marisa_trie trie, const char* file) {
    try {
        static_cast<const marisa::Trie*>(trie)->save(file);
        return ::MARISA_OK;
    } catch (const marisa::Exception &ex) {
        return ex.error_code();
    }
}

marisa_result marisa_trie_lookup(const marisa_trie trie, marisa_agent agent) {
    marisa_result result = {::MARISA_OK, false};
    try {
        marisa::Agent& ma = *static_cast<marisa::Agent*>(agent);
        if (static_cast<const marisa::Trie*>(trie)->lookup(ma)) {
            result.found = true;
            result.id = ma.key().id();
            result.str = ma.key().ptr();
            result.len = ma.key().length();
        }
    } catch (const marisa::Exception &ex) {
        result.err = ex.error_code();
    }
    return result;
}


marisa_result marisa_trie_reverse_lookup(const marisa_trie trie, marisa_agent agent){
    marisa_result result = {::MARISA_OK, false};
    try {
        marisa::Agent& ma = *static_cast<marisa::Agent*>(agent);
        static_cast<const marisa::Trie*>(trie)->reverse_lookup(ma);
        result.found = true;
        result.id = ma.key().id();
        result.str = ma.key().ptr();
        result.len = ma.key().length();
    } catch (const marisa::Exception &ex) {
        result.err = ex.error_code();
    }
    return result;
}

marisa_result marisa_trie_common_prefix_search(const marisa_trie trie, marisa_agent agent) {
    marisa_result result = {::MARISA_OK, false};
    try {
        marisa::Agent& ma = *static_cast<marisa::Agent*>(agent);
        if (static_cast<const marisa::Trie*>(trie)->common_prefix_search(ma)) {
            result.found = true;
            result.id = ma.key().id();
            result.str = ma.key().ptr();
            result.len = ma.key().length();
        }
    } catch (const marisa::Exception &ex) {
        result.err = ex.error_code();
    }
    return result;
}

marisa_result marisa_trie_predictive_search(const marisa_trie trie, marisa_agent agent) {
    marisa_result result = {::MARISA_OK, false};
    try {
        marisa::Agent& ma = *static_cast<marisa::Agent*>(agent);
        if (static_cast<const marisa::Trie*>(trie)->predictive_search(ma)) {
            result.found = true;
            result.id = ma.key().id();
            result.str = ma.key().ptr();
            result.len = ma.key().length();
        }
    } catch (const marisa::Exception &ex) {
        result.err = ex.error_code();
    }
    return result;
}



marisa_num marisa_trie_num_tries(const marisa_trie trie) {
    marisa_num ret = {::MARISA_OK};
    try {
        ret.num = static_cast<const marisa::Trie*>(trie)->num_tries();
    } catch (const marisa::Exception &ex) {
        ret.err = ex.error_code();
    }
    return ret;
}

marisa_num marisa_trie_num_keys(const marisa_trie trie) {
    marisa_num ret = {::MARISA_OK};
    try {
        ret.num = static_cast<const marisa::Trie*>(trie)->num_keys();
    } catch (const marisa::Exception &ex) {
        ret.err = ex.error_code();
    }
    return ret;
}

marisa_num marisa_trie_num_nodes(const marisa_trie trie) {
    marisa_num ret = {::MARISA_OK};
    try {
        ret.num = static_cast<const marisa::Trie*>(trie)->num_nodes();
    } catch (const marisa::Exception &ex) {
        ret.err = ex.error_code();
    }
    return ret;
}

marisa_num marisa_trie_size(const marisa_trie trie) {
    marisa_num ret = {::MARISA_OK};
    try {
        ret.num = static_cast<const marisa::Trie*>(trie)->size();
    } catch (const marisa::Exception &ex) {
        ret.err = ex.error_code();
    }
    return ret;
}

marisa_num marisa_trie_total_size(const marisa_trie trie) {
    marisa_num ret = {::MARISA_OK};
    try {
        ret.num = static_cast<const marisa::Trie*>(trie)->total_size();
    } catch (const marisa::Exception &ex) {
        ret.err = ex.error_code();
    }
    return ret;
}

marisa_num marisa_trie_io_size(const marisa_trie trie) {
    marisa_num ret = {::MARISA_OK};
    try {
        ret.num = static_cast<const marisa::Trie*>(trie)->io_size();
    } catch (const marisa::Exception &ex) {
        ret.err = ex.error_code();
    }
    return ret;
}
